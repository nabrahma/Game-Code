package repository

import (
    "context"

    "github.com/google/uuid"
    "github.com/gc-platform/api/internal/domain"
    "github.com/gc-platform/api/pkg/pagination"
    "gorm.io/gorm"
)

type ProblemRepo interface {
    List(ctx context.Context, filter domain.ProblemFilter, p pagination.Params) (*pagination.Page[domain.ProblemSummary], error)
    GetBySlug(ctx context.Context, slug string) (*domain.Problem, error)
    ToggleFavorite(ctx context.Context, userID uuid.UUID, problemID uuid.UUID) (bool, error)
}

type problemRepo struct {
    db *gorm.DB
}

func NewProblemRepo(db *gorm.DB) ProblemRepo {
    return &problemRepo{db: db}
}

func (r *problemRepo) List(ctx context.Context, filter domain.ProblemFilter, p pagination.Params) (*pagination.Page[domain.ProblemSummary], error) {
    // Stub implementation
    problems := []domain.ProblemSummary{
        {ID: uuid.New(), Slug: "two-sum", Title: "Two Sum", Difficulty: domain.DifficultyEasy, AcceptanceRate: 85.5},
        {ID: uuid.New(), Slug: "a-star-pathfinding", Title: "A* Pathfinding on Grid", Difficulty: domain.DifficultyHard, AcceptanceRate: 32.1},
        {ID: uuid.New(), Slug: "inventory-system", Title: "Inventory System Array", Difficulty: domain.DifficultyMedium, AcceptanceRate: 65.0},
        {ID: uuid.New(), Slug: "dialogue-tree", Title: "Traverse Dialogue Tree", Difficulty: domain.DifficultyMedium, AcceptanceRate: 50.2},
        {ID: uuid.New(), Slug: "collision-detection", Title: "AABB Collision Detection", Difficulty: domain.DifficultyEasy, AcceptanceRate: 90.0},
    }

    if filter.Difficulty != nil && *filter.Difficulty != "" {
        var filtered []domain.ProblemSummary
        for _, prob := range problems {
            if prob.Difficulty == domain.Difficulty(*filter.Difficulty) {
                filtered = append(filtered, prob)
            }
        }
        problems = filtered
    }

    return &pagination.Page[domain.ProblemSummary]{
        Items:      problems,
        Total:      int64(len(problems)),
        Page:       p.Page,
        Size:       p.Size,
        TotalPages: 1,
    }, nil
}

func (r *problemRepo) GetBySlug(ctx context.Context, slug string) (*domain.Problem, error) {
    // Stub implementation
    return &domain.Problem{
        Slug:           slug,
        Title:          "Sample Problem Title",
        Difficulty:     domain.DifficultyMedium,
        Status:         domain.ProblemStatusPublished,
        Description:    "This is a sample description of the problem context. You need to implement an algorithm to solve this.\n\n### Requirements\n- Be fast\n- Be memory efficient",
        Constraints:    "1 <= N <= 100",
        AcceptanceRate: 50.0,
        Examples: []domain.ProblemExample{
            {ID: uuid.New(), OrderIndex: 1, Input: "nums = [1,2,3]", Output: "[1,2,3]", Explanation: "Just return the array."},
        },
        Hints: []domain.ProblemHint{
            {ID: uuid.New(), OrderIndex: 1, Content: "Think about using a hash map."},
        },
        StarterCode: []domain.StarterCode{
            {Language: domain.LanguageCSharp, Code: "public class Solution {\n    public int[] Solve(int[] nums) {\n        \n    }\n}"},
            {Language: domain.LanguageCPP, Code: "class Solution {\npublic:\n    vector<int> solve(vector<int>& nums) {\n        \n    }\n};"},
        },
    }, nil
}

func (r *problemRepo) ToggleFavorite(ctx context.Context, userID uuid.UUID, problemID uuid.UUID) (bool, error) {
    var count int64
    r.db.WithContext(ctx).Model(&domain.Favorite{}).
        Where("user_id = ? AND problem_id = ?", userID, problemID).
        Count(&count)

    if count > 0 {
        // Already favorited, so unfavorite
        err := r.db.WithContext(ctx).
            Where("user_id = ? AND problem_id = ?", userID, problemID).
            Delete(&domain.Favorite{}).Error
        return false, err
    } else {
        // Not favorited, so add favorite
        err := r.db.WithContext(ctx).Create(&domain.Favorite{
            UserID:    userID,
            ProblemID: problemID,
        }).Error
        return true, err
    }
}
