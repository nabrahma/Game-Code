package service

import (
    "context"

    "encoding/json"
    "time"

    "github.com/google/uuid"
    "github.com/gc-platform/api/internal/cache"
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
    repo  repository.ListRepo
    cache cache.Cache
}

func NewListService(repo repository.ListRepo, c cache.Cache) ListService {
    return &listService{repo: repo, cache: c}
}

func (s *listService) GetCuratedLists(ctx context.Context) ([]domain.ProblemList, error) {
    cacheKey := "lists:curated"
    if cached, err := s.cache.Get(ctx, cacheKey); err == nil {
        var lists []domain.ProblemList
        if err := json.Unmarshal([]byte(cached), &lists); err == nil {
            return lists, nil
        }
    }

    lists, err := s.repo.ListCurated(ctx)
    if err != nil {
        return nil, err
    }

    if cacheBytes, err := json.Marshal(lists); err == nil {
        _ = s.cache.Set(ctx, cacheKey, string(cacheBytes), 15*time.Minute)
    }
    return lists, nil
}

func (s *listService) GetUserLists(ctx context.Context, userID uuid.UUID) ([]domain.ProblemList, error) {
    return s.repo.ListByUser(ctx, userID)
}

func (s *listService) GetListBySlug(ctx context.Context, slug string) (*domain.ProblemList, error) {
    cacheKey := "list:detail:" + slug
    if cached, err := s.cache.Get(ctx, cacheKey); err == nil {
        var list domain.ProblemList
        if err := json.Unmarshal([]byte(cached), &list); err == nil {
            return &list, nil
        }
    }

    list, err := s.repo.GetBySlug(ctx, slug)
    if err != nil {
        return nil, err
    }

    if cacheBytes, err := json.Marshal(list); err == nil {
        _ = s.cache.Set(ctx, cacheKey, string(cacheBytes), 5*time.Minute)
    }
    return list, nil
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
