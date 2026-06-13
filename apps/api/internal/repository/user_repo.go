package repository

import (
    "context"

    "github.com/google/uuid"
    "github.com/gc-platform/api/internal/domain"
    "gorm.io/gorm"
    "gorm.io/gorm/clause"
)

type UserRepo interface {
    ToggleFavorite(ctx context.Context, userID uuid.UUID, slug string) error
    UpsertProgress(ctx context.Context, userID uuid.UUID, problemID uuid.UUID, status string) error
}

type userRepo struct {
    db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
    return &userRepo{db: db}
}

func (r *userRepo) ToggleFavorite(ctx context.Context, userID uuid.UUID, slug string) error {
    // Requires problem ID, but signature only has slug. We can look it up first.
    // For now this is just a stub.
    return nil
}

func (r *userRepo) UpsertProgress(ctx context.Context, userID uuid.UUID, problemID uuid.UUID, status string) error {
    progress := domain.UserProblemProgress{
        UserID:    userID,
        ProblemID: problemID,
        Status:    status,
    }
    
    // GORM Upsert
    return r.db.WithContext(ctx).Clauses(clause.OnConflict{
        Columns:   []clause.Column{{Name: "user_id"}, {Name: "problem_id"}},
        DoUpdates: clause.AssignmentColumns([]string{"status", "attempt_count", "last_attempt"}),
    }).Create(&progress).Error
}
