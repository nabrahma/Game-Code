package domain

import (
    "time"
    "github.com/google/uuid"
)

type Difficulty string
const (
    DifficultyEasy   Difficulty = "easy"
    DifficultyMedium Difficulty = "medium"
    DifficultyHard   Difficulty = "hard"
)

type Language string
const (
    LanguageCSharp   Language = "csharp"
    LanguageCPP      Language = "cpp"
    LanguageLua      Language = "lua"
    LanguageGDScript Language = "gdscript"
)

type ProblemStatus string
const (
    ProblemStatusDraft     ProblemStatus = "draft"
    ProblemStatusReview    ProblemStatus = "review"
    ProblemStatusPublished ProblemStatus = "published"
    ProblemStatusArchived  ProblemStatus = "archived"
)

type Problem struct {
    ID                 uuid.UUID     `json:"id"`
    Slug               string        `json:"slug"`
    Title              string        `json:"title"`
    Difficulty         Difficulty    `json:"difficulty"`
    Status             ProblemStatus `json:"status"`
    Description        string        `json:"description"`
    Constraints        string        `json:"constraints"`
    GameEngineContext  string        `json:"game_engine_context,omitempty"`
    AcceptanceRate     float64       `json:"acceptance_rate"`
    SubmissionCount    int32         `json:"submission_count"`
    OrderIndex         int32         `json:"order_index"`
    CreatedAt          time.Time     `json:"created_at"`
    UpdatedAt          time.Time     `json:"updated_at"`
    PublishedAt        *time.Time    `json:"published_at,omitempty"`
    // Enriched fields (joined from related tables)
    Tags               []Tag         `json:"tags,omitempty"`
    Languages          []Language    `json:"languages,omitempty"`
    Examples           []ProblemExample `json:"examples,omitempty"`
    Hints              []ProblemHint    `json:"hints,omitempty"`
    StarterCode        []StarterCode    `json:"starter_code,omitempty"`
    VisibleTestCases   []TestCase       `json:"test_cases,omitempty"`
    Editorial          *Editorial       `json:"editorial,omitempty"`
    RelatedProblems    []ProblemSummary `json:"related_problems,omitempty"`
    // Per-user fields (nil when unauthenticated)
    UserStatus         *string          `json:"user_status,omitempty"`
    IsFavorite         *bool            `json:"is_favorite,omitempty"`
}


type ProblemSummary struct {
    ID             uuid.UUID  `json:"id"`
    Slug           string     `json:"slug"`
    Title          string     `json:"title"`
    Difficulty     Difficulty `json:"difficulty"`
    Tags           []Tag      `json:"tags"`
    Languages      []Language `json:"languages"`
    AcceptanceRate float64    `json:"acceptance_rate"`
    UserStatus     *string    `json:"user_status,omitempty"`
    IsFavorite     *bool      `json:"is_favorite,omitempty"`
}

type ProblemFilter struct {
    Difficulty *string
    Status     *string
    Search     *string
}

type Tag struct {
    ID       uuid.UUID `json:"id"`
    Name     string    `json:"name"`
    Slug     string    `json:"slug"`
    Category string    `json:"category,omitempty"`
}

type ProblemExample struct {
    ID          uuid.UUID `json:"id"`
    OrderIndex  int32     `json:"order_index"`
    Input       string    `json:"input"`
    Output      string    `json:"output"`
    Explanation string    `json:"explanation,omitempty"`
}

type ProblemHint struct {
    ID         uuid.UUID `json:"id"`
    OrderIndex int32     `json:"order_index"`
    Content    string    `json:"content"`
}

type TestCase struct {
    ID          uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
    ProblemID   uuid.UUID `json:"problem_id"`
    Input       string    `json:"input"`
    Output      string    `json:"output"`
    IsHidden    bool      `json:"is_hidden"`
    OrderIndex  int32     `json:"order_index"`
    TimeLimit   int32     `json:"time_limit"`
    MemoryLimit int32     `json:"memory_limit"`
}

type StarterCode struct {
    ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
    ProblemID uuid.UUID `json:"problem_id"`
    Language  Language  `json:"language"`
    Code      string    `json:"code"`
}

type Editorial struct {
    ID               uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
    ProblemID        uuid.UUID `json:"problem_id"`
    Content          string    `json:"content"`
    TimeComplexity   string    `json:"time_complexity,omitempty"`
    SpaceComplexity  string    `json:"space_complexity,omitempty"`
    UnityVariant     string    `json:"unity_variant,omitempty"`
    UnrealVariant    string    `json:"unreal_variant,omitempty"`
    GodotVariant     string    `json:"godot_variant,omitempty"`
}
