package repository

import (
    "context"

    "github.com/google/uuid"
    "github.com/gc-platform/api/internal/domain"
    "gorm.io/gorm"
)

type ListRepo interface {
    ListCurated(ctx context.Context) ([]domain.ProblemList, error)
    ListByUser(ctx context.Context, userID uuid.UUID) ([]domain.ProblemList, error)
    GetBySlug(ctx context.Context, slug string) (*domain.ProblemList, error)
    Create(ctx context.Context, list *domain.ProblemList) error
    AddProblem(ctx context.Context, listID, problemID uuid.UUID) error
    RemoveProblem(ctx context.Context, listID, problemID uuid.UUID) error
}

type listRepo struct {
    db *gorm.DB
}

func NewListRepo(db *gorm.DB) ListRepo {
    return &listRepo{db: db}
}

func (r *listRepo) ListCurated(ctx context.Context) ([]domain.ProblemList, error) {
    var lists []domain.ProblemList
    err := r.db.WithContext(ctx).
        Where("is_curated = ? AND is_public = ?", true, true).
        Order("created_at asc").
        Find(&lists).Error
    return lists, err
}

func (r *listRepo) ListByUser(ctx context.Context, userID uuid.UUID) ([]domain.ProblemList, error) {
    var lists []domain.ProblemList
    err := r.db.WithContext(ctx).
        Where("user_id = ?", userID).
        Order("created_at desc").
        Find(&lists).Error
    return lists, err
}

func (r *listRepo) GetBySlug(ctx context.Context, slug string) (*domain.ProblemList, error) {
    var list domain.ProblemList
    // Preload Items and their Problems
    err := r.db.WithContext(ctx).
        Preload("Items", func(db *gorm.DB) *gorm.DB {
            return db.Order("order_index asc")
        }).
        Preload("Items.Problem").
        Where("slug = ?", slug).
        First(&list).Error
    
    if err != nil {
        return nil, err
    }
    return &list, nil
}

func (r *listRepo) Create(ctx context.Context, list *domain.ProblemList) error {
    return r.db.WithContext(ctx).Create(list).Error
}

func (r *listRepo) AddProblem(ctx context.Context, listID, problemID uuid.UUID) error {
    var count int64
    r.db.WithContext(ctx).Model(&domain.ProblemListItem{}).Where("list_id = ?", listID).Count(&count)

    item := domain.ProblemListItem{
        ListID:     listID,
        ProblemID:  problemID,
        OrderIndex: int32(count),
    }
    return r.db.WithContext(ctx).Create(&item).Error
}

func (r *listRepo) RemoveProblem(ctx context.Context, listID, problemID uuid.UUID) error {
    return r.db.WithContext(ctx).
        Where("list_id = ? AND problem_id = ?", listID, problemID).
        Delete(&domain.ProblemListItem{}).Error
}
