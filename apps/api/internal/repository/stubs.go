package repository

import "github.com/gc-platform/api/internal/db/sqlc"

type UserRepo interface{}
type SubmissionRepo interface{}
type ListRepo interface{}

func NewUserRepo(q *sqlc.Queries) UserRepo { return nil }
func NewSubmissionRepo(q *sqlc.Queries) SubmissionRepo { return nil }
func NewListRepo(q *sqlc.Queries) ListRepo { return nil }
