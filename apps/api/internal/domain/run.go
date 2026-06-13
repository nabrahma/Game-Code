package domain

import (
    "time"
    "github.com/google/uuid"
)

type RunStatus string

const (
    RunStatusQueued  RunStatus = "queued"
    RunStatusRunning RunStatus = "running"
    RunStatusSuccess RunStatus = "success"
    RunStatusError   RunStatus = "error"
    RunStatusTimeout RunStatus = "timeout"
)

type RunRequest struct {
    ID        uuid.UUID `json:"id"`
    ProblemID uuid.UUID `json:"problem_id"`
    UserID    uuid.UUID `json:"user_id"`
    Language  Language  `json:"language"`
    Code      string    `json:"code"`
    Input     string    `json:"input,omitempty"` // For custom test cases
    CreatedAt time.Time `json:"created_at"`
}

type RunResult struct {
    ID         uuid.UUID `json:"id"`
    Status     RunStatus `json:"status"`
    Output     string    `json:"output"`      // stdout/stderr
    Error      string    `json:"error"`       // Compilation or system errors
    TimeMs     int64     `json:"time_ms"`     // Execution time in ms
    MemoryKb   int64     `json:"memory_kb"`   // Memory usage in KB
    FinishedAt time.Time `json:"finished_at"`
}

// Struct for Server-Sent Events (SSE) streaming
type RunStreamEvent struct {
    RunID  string `json:"run_id"`
    Status string `json:"status"`
    Output string `json:"output,omitempty"`
}
