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
    probHandler := NewProblemHandler(d.Prob, d.User)
    api.GET("/problems", probHandler.List)
    api.GET("/problems/:slug", probHandler.GetBySlug)
    api.POST("/problems/:slug/favorite", probHandler.ToggleFavorite)

    // Submissions
    // subHandler := NewSubmissionHandler(d.Sub)
    // api.POST("/submissions", subHandler.Create)
    
    // Execution
    runHandler := NewRunHandler(d.Run)
    api.POST("/run", runHandler.ExecuteCode)
    api.GET("/run/:runId/stream", runHandler.StreamRunLogs)
    
    // Lists
    listHandler := NewListHandler(d.List)
    api.GET("/lists/curated", listHandler.GetCuratedLists)
    api.GET("/lists/user", listHandler.GetUserLists)
    api.GET("/lists/:slug", listHandler.GetBySlug)
    api.POST("/lists", listHandler.CreateList)
    api.POST("/lists/:id/problems", listHandler.AddProblem)
    api.DELETE("/lists/:id/problems/:problemId", listHandler.RemoveProblem)
    
    // Auth (stubs for Phase 1)
    // authHandler := NewAuthHandler(d.Auth)
    // api.POST("/auth/logout", authHandler.Logout)
}
