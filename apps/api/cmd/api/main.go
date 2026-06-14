package main

import (
    "context"
    "log"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/labstack/echo/v4"
    "github.com/getsentry/sentry-go"
    "github.com/gc-platform/api/internal/config"
    "github.com/gc-platform/api/internal/db"
    "github.com/gc-platform/api/internal/handler"
    "github.com/gc-platform/api/internal/middleware"
    "github.com/gc-platform/api/internal/repository"
    "github.com/gc-platform/api/internal/service"
    "github.com/gc-platform/api/internal/cache"
    "github.com/redis/go-redis/v9"
    "go.uber.org/zap"
)

func main() {
    // 1. Load config
    cfg := config.Load()

    // 2. Init logger
    logger, _ := zap.NewProduction()
    defer logger.Sync()

    // Initialize Sentry
    if cfg.SentryDSN != "" {
        err := sentry.Init(sentry.ClientOptions{
            Dsn:              cfg.SentryDSN,
            Environment:      cfg.Environment,
            TracesSampleRate: 1.0,
        })
        if err != nil {
            logger.Error("Sentry initialization failed", zap.Error(err))
        } else {
            logger.Info("Sentry initialized")
            defer sentry.Flush(2 * time.Second)
        }
    }

    // 3. Connect to PostgreSQL via GORM
    gormDB, err := db.InitGORM(cfg)
    if err != nil {
        log.Fatal("cannot connect to db:", err)
    }

    // 4. Connect to Redis
    rdb := redis.NewClient(&redis.Options{Addr: cfg.RedisAddr, Password: cfg.RedisPassword})

    // 5. Build dependency tree
    cache     := cache.New(rdb)

    userRepo  := repository.NewUserRepo(gormDB)
    probRepo  := repository.NewProblemRepo(gormDB)
    subRepo   := repository.NewSubmissionRepo(gormDB)
    listRepo  := repository.NewListRepo(gormDB)
    discussRepo := repository.NewDiscussRepo(gormDB)

    authSvc   := service.NewAuthService(userRepo, cache, cfg)
    probSvc   := service.NewProblemService(probRepo, cache)
    subSvc    := service.NewSubmissionService(subRepo, probRepo, userRepo, cache)
    runSvc    := service.NewRunService(cache, cfg)
    listSvc   := service.NewListService(listRepo, cache)
    discussSvc := service.NewDiscussService(discussRepo, probRepo)
    userSvc   := service.NewUserService(userRepo, subRepo)


    // 7. Build Echo router
    e := echo.New()
    e.HideBanner = true

    // 8. Register middleware
    if cfg.SentryDSN != "" {
        e.Use(middleware.SentryMiddleware())
    }
    middleware.Register(e, cfg, rdb, logger)

    // 9. Register routes
    handler.RegisterRoutes(e, &handler.Deps{
        Auth:   authSvc,
        Prob:   probSvc,
        Sub:    subSvc,
        Run:    runSvc,
        List:    listSvc,
        Discuss: discussSvc,
        User:    userSvc,
        Logger: logger,
        Config: cfg,
    })

    // 10. Graceful shutdown
    go func() {
        if err := e.Start(":" + cfg.Port); err != nil {
            logger.Info("shutting down server", zap.Error(err))
        }
    }()

    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit

    ctx, cancel := context.WithTimeout(context.Background(), 10)
    defer cancel()
    e.Shutdown(ctx)
}
