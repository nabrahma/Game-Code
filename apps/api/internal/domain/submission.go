package domain

import (
    "time"
    "github.com/google/uuid"
)

type SubmissionVerdict string
const (
    VerdictPending            SubmissionVerdict = "pending"
    VerdictAccepted           SubmissionVerdict = "accepted"
    VerdictWrongAnswer        SubmissionVerdict = "wrong_answer"
    VerdictTimeLimitExceeded  SubmissionVerdict = "time_limit_exceeded"
    VerdictMemoryLimitExceeded SubmissionVerdict = "memory_limit_exceeded"
    VerdictRuntimeError       SubmissionVerdict = "runtime_error"
    VerdictCompileError       SubmissionVerdict = "compile_error"
    VerdictInternalError      SubmissionVerdict = "internal_error"
)

type Submission struct {
    ID               uuid.UUID         `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
    UserID           uuid.UUID         `json:"user_id" gorm:"type:uuid;not null"`
    ProblemID        uuid.UUID         `json:"problem_id" gorm:"type:uuid;not null"`
    Language         Language          `json:"language" gorm:"type:language;not null"`
    Code             string            `json:"code" gorm:"not null"`
    Verdict          SubmissionVerdict `json:"verdict" gorm:"type:submission_verdict;default:'pending';not null"`
    RuntimeMs        *int32            `json:"runtime_ms,omitempty"`
    MemoryKb         *int32            `json:"memory_kb,omitempty"`
    ErrorMessage     *string           `json:"error_message,omitempty"`
    PassedTestCount  *int32            `json:"passed_test_count,omitempty"`
    TotalTestCount   *int32            `json:"total_test_count,omitempty"`
    Results          []TestResult      `json:"results,omitempty" gorm:"foreignKey:SubmissionID"`
    CreatedAt        time.Time         `json:"created_at" gorm:"default:now();not null"`
    UpdatedAt        time.Time         `json:"updated_at" gorm:"default:now();not null"`
}

func (Submission) TableName() string {
    return "submissions"
}

type TestResult struct {
    ID           uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
    SubmissionID uuid.UUID `json:"submission_id" gorm:"type:uuid;not null"`
    Input        string    `json:"input" gorm:"not null"`
    Expected     string    `json:"expected" gorm:"not null"`
    Actual       string    `json:"actual" gorm:"not null;default:''"`
    Passed       bool      `json:"passed" gorm:"not null;default:false"`
    RuntimeMs    *int32    `json:"runtime_ms,omitempty"`
    MemoryKb     *int32    `json:"memory_kb,omitempty"`
}

func (TestResult) TableName() string {
    return "submission_test_results"
}
