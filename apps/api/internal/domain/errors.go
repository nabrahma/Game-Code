package domain

import "errors"

var (
    ErrNotFound         = errors.New("not found")
    ErrUnauthorized     = errors.New("unauthorized")
    ErrForbidden        = errors.New("forbidden")
    ErrConflict         = errors.New("conflict")
    ErrValidation       = errors.New("validation error")
    ErrRateLimit        = errors.New("rate limit exceeded")
    ErrInternalExec     = errors.New("execution internal error")
)
