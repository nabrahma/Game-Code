package service

import (
    "context"
    "encoding/json"
    "fmt"
    "time"

    "github.com/gc-platform/api/internal/cache"
    "github.com/gc-platform/api/internal/domain"
    "github.com/gc-platform/api/internal/repository"
)

type ProblemService interface {
    ListProblems(ctx context.Context, filter repository.ProblemFilter) ([]domain.ProblemSummary, int, error)
    GetProblem(ctx context.Context, slug string) (*domain.Problem, error)
}

type problemService struct {
    repo  repository.ProblemRepository
    cache cache.Cache
}

func NewProblemService(repo repository.ProblemRepository, c cache.Cache) ProblemService {
    return &problemService{repo: repo, cache: c}
}

type listCacheResult struct {
    Problems []domain.ProblemSummary `json:"problems"`
    Total    int                     `json:"total"`
}

func (s *problemService) ListProblems(ctx context.Context, filter repository.ProblemFilter) ([]domain.ProblemSummary, int, error) {
    diff := ""
    if filter.Difficulty != nil {
        diff = *filter.Difficulty
    }
    search := ""
    if filter.Search != nil {
        search = *filter.Search
    }

    cacheKey := fmt.Sprintf("problems:list:%s:%s:%s:%d:%d", diff, search, filter.Sort, filter.Offset, filter.Limit)
    
    // Try cache
    val, err := s.cache.Get(ctx, cacheKey)
    if err == nil && val != "" {
        var res listCacheResult
        if json.Unmarshal([]byte(val), &res) == nil {
            return res.Problems, res.Total, nil
        }
    }

    // Fetch from DB
    problems, total, err := s.repo.List(ctx, filter)
    if err != nil {
        return nil, 0, err
    }

    // Set cache (30s TTL per PRD)
    res := listCacheResult{Problems: problems, Total: total}
    if bytes, err := json.Marshal(res); err == nil {
        _ = s.cache.Set(ctx, cacheKey, string(bytes), 30*time.Second)
    }

    return problems, total, nil
}

func (s *problemService) GetProblem(ctx context.Context, slug string) (*domain.Problem, error) {
    return s.repo.GetBySlug(ctx, slug)
}
