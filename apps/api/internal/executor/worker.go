package executor

import (
    "github.com/gc-platform/api/internal/config"
    "github.com/gc-platform/api/internal/repository"
    "github.com/redis/go-redis/v9"
    "go.uber.org/zap"
)

type Worker struct {
    cfg      *config.Config
    rdb      *redis.Client
    subRepo  repository.SubmissionRepo
    probRepo repository.ProblemRepo
    logger   *zap.Logger
}

func NewWorker(cfg *config.Config, rdb *redis.Client, subRepo repository.SubmissionRepo, probRepo repository.ProblemRepo, logger *zap.Logger) *Worker {
    return &Worker{
        cfg:      cfg,
        rdb:      rdb,
        subRepo:  subRepo,
        probRepo: probRepo,
        logger:   logger,
    }
}

func (w *Worker) Start() {
    w.logger.Info("Starting execution worker...")
    // In a real implementation, this would start the asynq server
    // and bind to the redis client to listen for execution tasks.
}
