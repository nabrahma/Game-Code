package handler

import (
    "github.com/labstack/echo/v4"
    "github.com/gc-platform/api/internal/config"
    "github.com/gc-platform/api/internal/service"
    "go.uber.org/zap"
)

type Deps struct {
    Auth   service.AuthService
    Prob   service.ProblemService
    Sub    service.SubmissionService
    Run    service.RunService
    List   service.ListService
    User   service.UserService
    Logger *zap.Logger
    Config *config.Config
}

func RegisterRoutes(e *echo.Echo, d *Deps) {
    // API Group
    api := e.Group("/api")

    // Problems
    probHandler := NewProblemHandler(d.Prob)
    api.GET("/problems", probHandler.List)
    api.GET("/problems/:slug", probHandler.GetBySlug)

    // Submissions
    // subHandler := NewSubmissionHandler(d.Sub)
    // api.POST("/submissions", subHandler.Create)
    
    // Auth (stubs for Phase 1)
    // authHandler := NewAuthHandler(d.Auth)
    // api.POST("/auth/logout", authHandler.Logout)
}
