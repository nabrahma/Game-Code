package repository

import (
    "context"

    "github.com/google/uuid"
    "github.com/gc-platform/api/internal/domain"
    "gorm.io/gorm"
)

type SubmissionRepo interface {
    Create(ctx context.Context, sub *domain.Submission) error
    UpdateVerdict(ctx context.Context, id uuid.UUID, verdict domain.SubmissionVerdict, runtime int32, memory int32, errorMsg string, passed int32, total int32) error
    GetByID(ctx context.Context, id uuid.UUID) (*domain.Submission, error)
    ListByUser(ctx context.Context, userID uuid.UUID) ([]domain.Submission, error)
}

type submissionRepo struct {
    db *gorm.DB
}

func NewSubmissionRepo(db *gorm.DB) SubmissionRepo {
    return &submissionRepo{db: db}
}

func (r *submissionRepo) Create(ctx context.Context, sub *domain.Submission) error {
    return r.db.WithContext(ctx).Create(sub).Error
}

func (r *submissionRepo) UpdateVerdict(ctx context.Context, id uuid.UUID, verdict domain.SubmissionVerdict, runtime int32, memory int32, errorMsg string, passed int32, total int32) error {
    updates := map[string]interface{}{
        "verdict":           verdict,
        "runtime_ms":        runtime,
        "memory_kb":         memory,
        "error_message":     errorMsg,
        "passed_test_count": passed,
        "total_test_count":  total,
    }
    return r.db.WithContext(ctx).Model(&domain.Submission{}).Where("id = ?", id).Updates(updates).Error
}

func (r *submissionRepo) GetByID(ctx context.Context, id uuid.UUID) (*domain.Submission, error) {
    var sub domain.Submission
    err := r.db.WithContext(ctx).Preload("Results").First(&sub, "id = ?", id).Error
    if err != nil {
        return nil, err
    }
    return &sub, nil
}

func (r *submissionRepo) ListByUser(ctx context.Context, userID uuid.UUID) ([]domain.Submission, error) {
    var subs []domain.Submission
    err := r.db.WithContext(ctx).Where("user_id = ?", userID).Order("created_at desc").Limit(50).Find(&subs).Error
    return subs, err
}
