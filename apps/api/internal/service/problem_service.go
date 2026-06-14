package service

import (
    "context"
    "encoding/json"
    "fmt"
    "time"

    "github.com/google/uuid"
    "github.com/gc-platform/api/internal/cache"
    "github.com/gc-platform/api/internal/domain"
    "github.com/gc-platform/api/internal/repository"
    "github.com/gc-platform/api/pkg/pagination"
)

type ProblemService interface {
    GetProblems(ctx context.Context, filter domain.ProblemFilter, p pagination.Params) (*pagination.Page[domain.ProblemSummary], error)
    GetProblem(ctx context.Context, slug string) (*domain.Problem, error)
    ToggleFavorite(ctx context.Context, userID uuid.UUID, problemID uuid.UUID) (bool, error)

    // Admin CMS methods
    CreateProblem(ctx context.Context, p *domain.Problem) error
    UpdateProblem(ctx context.Context, slug string, updates map[string]interface{}) error
    DeleteProblem(ctx context.Context, id uuid.UUID) error
    UpsertTestCase(ctx context.Context, tc *domain.TestCase) error
    DeleteTestCase(ctx context.Context, id uuid.UUID) error
    UpsertStarterCode(ctx context.Context, sc *domain.StarterCode) error
    UpsertEditorial(ctx context.Context, e *domain.Editorial) error
}

type problemService struct {
    repo  repository.ProblemRepo
    cache cache.Cache
}

func NewProblemService(repo repository.ProblemRepo, c cache.Cache) ProblemService {
    return &problemService{
        repo:  repo,
        cache: c,
    }
}

func (s *problemService) GetProblems(ctx context.Context, filter domain.ProblemFilter, p pagination.Params) (*pagination.Page[domain.ProblemSummary], error) {
    diff := ""
    if filter.Difficulty != nil {
        diff = *filter.Difficulty
    }
    
    cacheKey := fmt.Sprintf("problems:list:%s:%d:%d", diff, p.Page, p.Size)
    
    if cached, err := s.cache.Get(ctx, cacheKey); err == nil {
        var result pagination.Page[domain.ProblemSummary]
        if err := json.Unmarshal([]byte(cached), &result); err == nil {
            return &result, nil
        }
    }

    result, err := s.repo.List(ctx, filter, p)
    if err != nil {
        return nil, err
    }

    if cacheBytes, err := json.Marshal(result); err == nil {
        _ = s.cache.Set(ctx, cacheKey, string(cacheBytes), 5*time.Minute)
    }

    return result, nil
}

func (s *problemService) GetProblem(ctx context.Context, slug string) (*domain.Problem, error) {
    cacheKey := "problem:detail:" + slug
    
    if cached, err := s.cache.Get(ctx, cacheKey); err == nil {
        var prob domain.Problem
        if err := json.Unmarshal([]byte(cached), &prob); err == nil {
            return &prob, nil
        }
    }

    prob, err := s.repo.GetBySlug(ctx, slug)
    if err != nil {
        return nil, err
    }

    if cacheBytes, err := json.Marshal(prob); err == nil {
        _ = s.cache.Set(ctx, cacheKey, string(cacheBytes), 1*time.Hour)
    }

    return prob, nil
}

func (s *problemService) ToggleFavorite(ctx context.Context, userID uuid.UUID, problemID uuid.UUID) (bool, error) {
    return s.repo.ToggleFavorite(ctx, userID, problemID)
}

// Admin CMS Methods

func (s *problemService) CreateProblem(ctx context.Context, p *domain.Problem) error {
    return s.repo.CreateProblem(ctx, p)
}

func (s *problemService) UpdateProblem(ctx context.Context, slug string, updates map[string]interface{}) error {
    return s.repo.UpdateProblem(ctx, slug, updates)
}

func (s *problemService) DeleteProblem(ctx context.Context, id uuid.UUID) error {
    return s.repo.DeleteProblem(ctx, id)
}

func (s *problemService) UpsertTestCase(ctx context.Context, tc *domain.TestCase) error {
    return s.repo.UpsertTestCase(ctx, tc)
}

func (s *problemService) DeleteTestCase(ctx context.Context, id uuid.UUID) error {
    return s.repo.DeleteTestCase(ctx, id)
}

func (s *problemService) UpsertStarterCode(ctx context.Context, sc *domain.StarterCode) error {
    return s.repo.UpsertStarterCode(ctx, sc)
}

func (s *problemService) UpsertEditorial(ctx context.Context, e *domain.Editorial) error {
    return s.repo.UpsertEditorial(ctx, e)
}
