package db

import (
    "fmt"
    "log"

    "github.com/gc-platform/api/internal/config"
    "github.com/gc-platform/api/internal/domain"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

// InitGORM initializes and returns a GORM database connection
func InitGORM(cfg *config.Config) (*gorm.DB, error) {
    dsn := cfg.DatabaseURL
    if dsn == "" {
        // Fallback for local development
        dsn = "host=localhost user=postgres password=postgres dbname=gamecode port=5432 sslmode=disable"
    }

    // Configure GORM
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
        // Disable foreign key constraint when migrating (we use golang-migrate anyway)
        DisableForeignKeyConstraintWhenMigrating: true,
    })

    if err != nil {
        return nil, fmt.Errorf("failed to connect database: %w", err)
    }

    err = db.AutoMigrate(
        &domain.User{},
        &domain.Submission{},
        &domain.TestResult{},
        &domain.UserProblemProgress{},
    )
    if err != nil {
        log.Printf("GORM AutoMigrate failed (safe to ignore if using golang-migrate): %v", err)
    }

    log.Println("GORM connected to PostgreSQL successfully")
    return db, nil
}
