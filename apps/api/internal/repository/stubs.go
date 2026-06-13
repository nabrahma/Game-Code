package repository

import "github.com/gc-platform/api/internal/db/sqlc"

type UserRepo interface{
    ToggleFavorite(slug string) error
}
type SubmissionRepo interface{}
type ListRepo interface{}

func NewUserRepo(q *sqlc.Queries) UserRepo { return &userRepoStub{} }
func NewSubmissionRepo(q *sqlc.Queries) SubmissionRepo { return nil }
func NewListRepo(q *sqlc.Queries) ListRepo { return nil }

type userRepoStub struct{}

func (s *userRepoStub) ToggleFavorite(slug string) error {
    return nil // Optmistic mock
}
