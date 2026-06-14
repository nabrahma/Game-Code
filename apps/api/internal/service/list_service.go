package service

import (
    "context"

    "github.com/google/uuid"
    "github.com/gc-platform/api/internal/domain"
    "github.com/gc-platform/api/internal/repository"
)

type ListService interface {
    GetCuratedLists(ctx context.Context) ([]domain.ProblemList, error)
    GetUserLists(ctx context.Context, userID uuid.UUID) ([]domain.ProblemList, error)
    GetListBySlug(ctx context.Context, slug string) (*domain.ProblemList, error)
    CreateList(ctx context.Context, list *domain.ProblemList) error
    AddProblemToList(ctx context.Context, listID, problemID uuid.UUID) error
    RemoveProblemFromList(ctx context.Context, listID, problemID uuid.UUID) error
}

type listService struct {
    repo repository.ListRepo
}

func NewListService(repo repository.ListRepo) ListService {
    return &listService{repo: repo}
}

func (s *listService) GetCuratedLists(ctx context.Context) ([]domain.ProblemList, error) {
    return s.repo.ListCurated(ctx)
}

func (s *listService) GetUserLists(ctx context.Context, userID uuid.UUID) ([]domain.ProblemList, error) {
    return s.repo.ListByUser(ctx, userID)
}

func (s *listService) GetListBySlug(ctx context.Context, slug string) (*domain.ProblemList, error) {
    return s.repo.GetBySlug(ctx, slug)
}

func (s *listService) CreateList(ctx context.Context, list *domain.ProblemList) error {
    return s.repo.Create(ctx, list)
}

func (s *listService) AddProblemToList(ctx context.Context, listID, problemID uuid.UUID) error {
    return s.repo.AddProblem(ctx, listID, problemID)
}

func (s *listService) RemoveProblemFromList(ctx context.Context, listID, problemID uuid.UUID) error {
    return s.repo.RemoveProblem(ctx, listID, problemID)
}
