package main

import (
    "context"
    "log"
    "os"
    "os/signal"
    "syscall"

    "github.com/labstack/echo/v4"
    "github.com/gc-platform/api/internal/config"
    "github.com/gc-platform/api/internal/db"
    "github.com/gc-platform/api/internal/executor"
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

    authSvc   := service.NewAuthService(userRepo, cache, cfg)
    probSvc   := service.NewProblemService(probRepo, cache)
    subSvc    := service.NewSubmissionService(subRepo, probRepo, userRepo, cache)
    runSvc    := service.NewRunService(cache, cfg)
    listSvc   := service.NewListService(listRepo)
    userSvc   := service.NewUserService(userRepo, subRepo)

    // 6. Start asynq workers (goroutine-based, same process as Option A)
    execWorker := executor.NewWorker(cfg, rdb, subRepo, probRepo, logger)
    go execWorker.Start()

    // 7. Build Echo router
    e := echo.New()
    e.HideBanner = true

    // 8. Register middleware
    middleware.Register(e, cfg, rdb, logger)

    // 9. Register routes
    handler.RegisterRoutes(e, &handler.Deps{
        Auth:   authSvc,
        Prob:   probSvc,
        Sub:    subSvc,
        Run:    runSvc,
        List:   listSvc,
        User:   userSvc,
        Logger: logger,
        Config: cfg,
    })

    // 10. Graceful shutdown
    go func() {
        if err := e.Start(":" + cfg.Port); err != nil {
            logger.Info("shutting down server")
        }
    }()

    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit

    ctx, cancel := context.WithTimeout(context.Background(), 10)
    defer cancel()
    e.Shutdown(ctx)
}
