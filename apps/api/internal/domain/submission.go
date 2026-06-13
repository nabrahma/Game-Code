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
    ID               uuid.UUID         `json:"id"`
    UserID           uuid.UUID         `json:"user_id"`
    ProblemID        uuid.UUID         `json:"problem_id"`
    ProblemSlug      string            `json:"problem_slug"`
    ProblemTitle     string            `json:"problem_title"`
    Language         Language          `json:"language"`
    Code             string            `json:"code"`
    Verdict          SubmissionVerdict `json:"verdict"`
    RuntimeMs        *int32            `json:"runtime_ms,omitempty"`
    MemoryKb         *int32            `json:"memory_kb,omitempty"`
    ErrorMessage     *string           `json:"error_message,omitempty"`
    PassedTestCount  *int32            `json:"passed_test_count,omitempty"`
    TotalTestCount   *int32            `json:"total_test_count,omitempty"`
    Results          []TestResult      `json:"results,omitempty"`
    CreatedAt        time.Time         `json:"created_at"`
}

type TestResult struct {
    ID         uuid.UUID `json:"id"`
    Input      string    `json:"input"`
    Expected   string    `json:"expected"`
    Actual     string    `json:"actual"`
    Passed     bool      `json:"passed"`
    RuntimeMs  *int32    `json:"runtime_ms,omitempty"`
    MemoryKb   *int32    `json:"memory_kb,omitempty"`
}
