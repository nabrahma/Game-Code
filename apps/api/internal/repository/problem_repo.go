package repository

import (
    "context"

    "github.com/google/uuid"
    "github.com/gc-platform/api/internal/domain"
    "github.com/gc-platform/api/pkg/pagination"
    "gorm.io/gorm"
    "gorm.io/gorm/clause"
)

type ProblemRepo interface {
    List(ctx context.Context, filter domain.ProblemFilter, p pagination.Params) (*pagination.Page[domain.ProblemSummary], error)
    GetBySlug(ctx context.Context, slug string) (*domain.Problem, error)
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
        Title:          "Two Sum",
        Difficulty:     domain.DifficultyEasy,
        Status:         domain.ProblemStatusPublished,
        Description:    "Given an array of integers `nums` and an integer `target`, return indices of the two numbers such that they add up to `target`.\n\nYou may assume that each input would have **exactly one solution**, and you may not use the same element twice.\n\nYou can return the answer in any order.",
        Constraints:    "2 <= nums.length <= 10^4\n-10^9 <= nums[i] <= 10^9\n-10^9 <= target <= 10^9\nOnly one valid answer exists.",
        AcceptanceRate: 49.2,
        Examples: []domain.ProblemExample{
            {ID: uuid.New(), OrderIndex: 1, Input: "nums = [2,7,11,15], target = 9", Output: "[0,1]", Explanation: "Because nums[0] + nums[1] == 9, we return [0, 1]."},
            {ID: uuid.New(), OrderIndex: 2, Input: "nums = [3,2,4], target = 6", Output: "[1,2]", Explanation: ""},
        },
        Hints: []domain.ProblemHint{
            {ID: uuid.New(), OrderIndex: 1, Content: "A really brute force way would be to search for all possible pairs of numbers but that would be too slow. Again, it's best to try out brute force solutions for just for completeness. It is from these brute force solutions that you can come up with optimizations."},
            {ID: uuid.New(), OrderIndex: 2, Content: "So, if we fix one of the numbers, say `x`, we have to scan the entire array to find the next number `y` which is `value - x` where value is the input parameter. Can we change our array keeping so that this search becomes faster?"},
        },
        StarterCode: []domain.StarterCode{
            {Language: domain.LanguageCSharp, Code: "public class Solution {\n    public int[] Solve(int[] nums) {\n        \n    }\n}"},
            {Language: domain.LanguageCPP, Code: "class Solution {\npublic:\n    vector<int> solve(vector<int>& nums) {\n        \n    }\n};"},
            {Language: domain.LanguageLua, Code: "local Solution = {}\n\nfunction Solution:solve(nums)\n    \nend\n\nreturn Solution"},
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

// Admin CMS Methods

func (r *problemRepo) CreateProblem(ctx context.Context, p *domain.Problem) error {
    return r.db.WithContext(ctx).Create(p).Error
}

func (r *problemRepo) UpdateProblem(ctx context.Context, slug string, updates map[string]interface{}) error {
    return r.db.WithContext(ctx).Model(&domain.Problem{}).Where("slug = ?", slug).Updates(updates).Error
}

func (r *problemRepo) DeleteProblem(ctx context.Context, id uuid.UUID) error {
    return r.db.WithContext(ctx).Delete(&domain.Problem{}, id).Error
}

func (r *problemRepo) UpsertTestCase(ctx context.Context, tc *domain.TestCase) error {
    if tc.ID == uuid.Nil {
        tc.ID = uuid.New()
        return r.db.WithContext(ctx).Table("test_cases").Create(tc).Error
    }
    return r.db.WithContext(ctx).Table("test_cases").Save(tc).Error
}

func (r *problemRepo) DeleteTestCase(ctx context.Context, id uuid.UUID) error {
    return r.db.WithContext(ctx).Table("test_cases").Where("id = ?", id).Delete(nil).Error
}

func (r *problemRepo) UpsertStarterCode(ctx context.Context, sc *domain.StarterCode) error {
    if sc.ID == uuid.Nil {
        sc.ID = uuid.New()
    }
    return r.db.WithContext(ctx).Table("starter_code").
        Clauses(clause.OnConflict{
            Columns:   []clause.Column{{Name: "problem_id"}, {Name: "language"}},
            DoUpdates: clause.AssignmentColumns([]string{"code"}),
        }).Create(sc).Error
}

func (r *problemRepo) UpsertEditorial(ctx context.Context, e *domain.Editorial) error {
    if e.ID == uuid.Nil {
        e.ID = uuid.New()
    }
    return r.db.WithContext(ctx).Table("editorials").
        Clauses(clause.OnConflict{
            Columns:   []clause.Column{{Name: "problem_id"}},
            DoUpdates: clause.AssignmentColumns([]string{"content", "time_complexity", "space_complexity", "unity_variant", "unreal_variant", "godot_variant", "updated_at"}),
        }).Create(e).Error
}

