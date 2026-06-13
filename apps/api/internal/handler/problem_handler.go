package handler

import (
    "net/http"
    "strconv"

    "github.com/labstack/echo/v4"
    "github.com/gc-platform/api/internal/repository"
    "github.com/gc-platform/api/internal/service"
)

type ProblemHandler struct {
    probSvc service.ProblemService
    userSvc service.UserService
}

func NewProblemHandler(probSvc service.ProblemService, userSvc service.UserService) *ProblemHandler {
    return &ProblemHandler{probSvc: probSvc, userSvc: userSvc}
}

func (h *ProblemHandler) List(c echo.Context) error {
    filter := repository.ProblemFilter{
        Sort:   c.QueryParam("sort"),
        Limit:  10,
        Offset: 0,
    }
    
    if diff := c.QueryParam("difficulty"); diff != "" {
        filter.Difficulty = &diff
    }
    if search := c.QueryParam("q"); search != "" {
        filter.Search = &search
    }
    if limitStr := c.QueryParam("limit"); limitStr != "" {
        if l, err := strconv.Atoi(limitStr); err == nil && l > 0 && l <= 100 {
            filter.Limit = int32(l)
        }
    }
    if offsetStr := c.QueryParam("offset"); offsetStr != "" {
        if o, err := strconv.Atoi(offsetStr); err == nil && o >= 0 {
            filter.Offset = int32(o)
        }
    }

    problems, total, err := h.probSvc.ListProblems(c.Request().Context(), filter)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "problems": problems,
        "total":    total,
    })
}

func (h *ProblemHandler) GetBySlug(c echo.Context) error {
    slug := c.Param("slug")
    if slug == "" {
        return echo.NewHTTPError(http.StatusBadRequest, "slug is required")
    }

    problem, err := h.probSvc.GetProblem(c.Request().Context(), slug)
    if err != nil {
        return echo.NewHTTPError(http.StatusNotFound, "problem not found")
    }

    return c.JSON(http.StatusOK, problem)
}

func (h *ProblemHandler) ToggleFavorite(c echo.Context) error {
    slug := c.Param("slug")
    if slug == "" {
        return echo.NewHTTPError(http.StatusBadRequest, "slug is required")
    }
    
    // In a real implementation we would extract UserID from context and call:
    // h.userSvc.ToggleFavorite(userID, slug)
    // For now we assume optimistic success
    return c.NoContent(http.StatusOK)
}
