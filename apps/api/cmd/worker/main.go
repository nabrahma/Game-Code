package main

import (
    "context"
    "log"
    "os"
    "os/signal"
    "syscall"

    "github.com/gc-platform/api/internal/config"
    "github.com/gc-platform/api/internal/db/sqlc"
    "github.com/gc-platform/api/internal/executor"
    "github.com/gc-platform/api/internal/repository"
    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/redis/go-redis/v9"
    "go.uber.org/zap"
)

func main() {
    cfg := config.Load()

    logger, _ := zap.NewProduction()
    defer logger.Sync()

    pool, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
    if err != nil {
        log.Fatal("cannot connect to db:", err)
    }
    defer pool.Close()

    rdb := redis.NewClient(&redis.Options{Addr: cfg.RedisAddr, Password: cfg.RedisPassword})
    queries := sqlc.New(pool)

    subRepo := repository.NewSubmissionRepo(queries)
    probRepo := repository.NewProblemRepo(queries)

    worker := executor.NewWorker(cfg, rdb, subRepo, probRepo, logger)
    go worker.Start()

    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit

    logger.Info("shutting down worker gracefully")
}
