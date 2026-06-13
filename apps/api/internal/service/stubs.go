package service

import (
    "github.com/gc-platform/api/internal/cache"
    "github.com/gc-platform/api/internal/config"
    "github.com/gc-platform/api/internal/repository"
)

type AuthService interface{}
type SubmissionService interface{}
type RunService interface{}
type ListService interface{}
type UserService interface{}

func NewAuthService(userRepo repository.UserRepo, cache cache.Cache, cfg *config.Config) AuthService { return nil }
func NewSubmissionService(subRepo repository.SubmissionRepo, probRepo repository.ProblemRepository, cache cache.Cache) SubmissionService { return nil }
func NewRunService(cache cache.Cache, cfg *config.Config) RunService { return nil }
func NewListService(listRepo repository.ListRepo) ListService { return nil }
func NewUserService(userRepo repository.UserRepo, subRepo repository.SubmissionRepo) UserService { return nil }
