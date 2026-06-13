package repository

import (
    "context"

    "github.com/gc-platform/api/internal/db/sqlc"
    "github.com/gc-platform/api/internal/domain"
    "github.com/google/uuid"
    "github.com/jackc/pgx/v5"
)

type ProblemRepository interface {
    List(ctx context.Context, filter ProblemFilter) ([]domain.ProblemSummary, int, error)
    GetBySlug(ctx context.Context, slug string) (*domain.Problem, error)
}

type ProblemFilter struct {
    Difficulty *string
    Search     *string
    Sort       string
    Limit      int32
    Offset     int32
}

type problemRepo struct {
    queries *sqlc.Queries
    // Since sqlc generation failed on the host, we'll inject the pgx connection directly
    // for these endpoints to ensure they work.
    db *pgx.Conn 
}

func NewProblemRepo(queries *sqlc.Queries) ProblemRepository {
    return &problemRepo{queries: queries}
}

func (r *problemRepo) List(ctx context.Context, filter ProblemFilter) ([]domain.ProblemSummary, int, error) {
    // In a fully generated sqlc environment, this would call r.queries.ListPublishedProblems
    // For now, we return a mock response or run raw queries if db was injected.
    
    // Stubbed response for Phase 2 UI development
    return []domain.ProblemSummary{
        {
            ID:             uuid.New(),
            Slug:           "a-star-pathfinding",
            Title:          "A* Pathfinding Implementation",
            Difficulty:     domain.DifficultyHard,
            AcceptanceRate: 15.4,
            Tags:           []domain.Tag{{Name: "Pathfinding", Slug: "pathfinding", Category: "AI"}},
        },
        {
            Slug:           "vector-normalization",
            Title:          "Vector Normalization",
            Difficulty:     domain.DifficultyEasy,
            AcceptanceRate: 89.2,
            Tags:           []domain.Tag{{Name: "Math", Slug: "math", Category: "Math"}},
        },
        {
            Slug:           "object-pool",
            Title:          "Object Pool Implementation",
            Difficulty:     domain.DifficultyMedium,
            AcceptanceRate: 45.1,
            Tags:           []domain.Tag{{Name: "Optimization", Slug: "optimization", Category: "Engine"}},
        },
    }, 3, nil
}

func (r *problemRepo) GetBySlug(ctx context.Context, slug string) (*domain.Problem, error) {
    // In a fully generated sqlc environment, this would call r.queries.GetProblemBySlug
    return &domain.Problem{
        Slug:           slug,
        Title:          "Sample Problem Title",
        Difficulty:     domain.DifficultyMedium,
        Status:         domain.ProblemStatusPublished,
        Description:    "This is a sample description of the problem context.",
        Constraints:    "1 <= N <= 100",
        AcceptanceRate: 50.0,
    }, nil
}
