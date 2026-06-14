package handler

import (
    "net/http"

    "github.com/labstack/echo/v4"
    "github.com/gc-platform/api/internal/config"
    "github.com/gc-platform/api/internal/service"
    "go.uber.org/zap"
)

type Deps struct {
    Auth    service.AuthService
    Prob    service.ProblemService
    Sub     service.SubmissionService
    Run     service.RunService
    List    service.ListService
    Discuss service.DiscussService
    User    service.UserService
    Logger  *zap.Logger
    Config  *config.Config
}

func RegisterRoutes(e *echo.Echo, d *Deps) {
    // Health Check
    e.GET("/health", func(c echo.Context) error {
        return c.JSON(http.StatusOK, map[string]string{"status": "ok", "version": "1.0"})
    })
    
    // API Group
    api := e.Group("/api")

    // Problems
    probHandler := NewProblemHandler(d.Prob, d.User)
    api.GET("/problems", probHandler.List)
    api.GET("/problems/:slug", probHandler.GetBySlug)
    api.POST("/problems/:slug/favorite", probHandler.ToggleFavorite)

    // Submissions
    subHandler := NewSubmissionHandler(d.Sub)
    api.POST("/submissions", subHandler.Create)
    
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
    
    // Discuss
    discussHandler := NewDiscussHandler(d.Discuss)
    api.GET("/problems/:slug/discuss", discussHandler.ListPostsByProblem)
    api.POST("/problems/:slug/discuss", discussHandler.CreatePost)

    api.GET("/discuss/:id", discussHandler.GetPostByID)
    api.GET("/discuss/:id/comments", discussHandler.ListCommentsForPost)
    api.POST("/discuss/:id/comments", discussHandler.CreateComment)
    api.POST("/discuss/:id/upvote", discussHandler.TogglePostUpvote)
    api.POST("/discuss/comments/:commentId/upvote", discussHandler.ToggleCommentUpvote)

    // Admin routes
    adminHandler := NewAdminHandler(d.Prob)
    admin := api.Group("/admin")
    admin.POST("/problems", adminHandler.CreateProblem)
    admin.PUT("/problems/:slug", adminHandler.UpdateProblem)
    admin.DELETE("/problems/:id", adminHandler.DeleteProblem)
    
    admin.POST("/testcases", adminHandler.UpsertTestCase)
    admin.DELETE("/testcases/:id", adminHandler.DeleteTestCase)
    
    admin.POST("/startercode", adminHandler.UpsertStarterCode)
    admin.POST("/editorials", adminHandler.UpsertEditorial)

    // Auth (stubs for Phase 1)
    // authHandler := NewAuthHandler(d.Auth)
    // api.POST("/auth/logout", authHandler.Logout)

    api.GET("/users/me", func(c echo.Context) error {
        return c.JSON(http.StatusOK, map[string]interface{}{
            "user": map[string]interface{}{
                "id": "00000000-0000-0000-0000-000000000001",
                "username": "demo_user",
                "email": "demo@example.com",
                "role": "user",
            },
        })
    })
}
