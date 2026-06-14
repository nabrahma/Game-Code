package service

import (
    "github.com/gc-platform/api/internal/cache"
    "github.com/gc-platform/api/internal/config"
    "github.com/gc-platform/api/internal/repository"
)

type AuthService interface{}
type UserService interface{}

func NewAuthService(u repository.UserRepo, c cache.Cache, cfg *config.Config) AuthService {
    return nil
}


func NewUserService(u repository.UserRepo, s repository.SubmissionRepo) UserService {
    return nil
}
