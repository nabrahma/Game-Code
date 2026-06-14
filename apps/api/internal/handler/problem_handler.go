package handler

import (
    "net/http"
    "strconv"

    "github.com/labstack/echo/v4"
    "github.com/gc-platform/api/internal/domain"
    "github.com/gc-platform/api/internal/service"
    "github.com/gc-platform/api/pkg/pagination"
    "github.com/google/uuid"
)

type ProblemHandler struct {
    probSvc service.ProblemService
    userSvc service.UserService
}

func NewProblemHandler(probSvc service.ProblemService, userSvc service.UserService) *ProblemHandler {
    return &ProblemHandler{probSvc: probSvc, userSvc: userSvc}
}

func (h *ProblemHandler) List(c echo.Context) error {
    filter := domain.ProblemFilter{}
    
    if diff := c.QueryParam("difficulty"); diff != "" {
        filter.Difficulty = &diff
    }
    if search := c.QueryParam("q"); search != "" {
        filter.Search = &search
    }

    params := pagination.Params{
        Page: 1,
        Size: 10,
    }

    if pageStr := c.QueryParam("page"); pageStr != "" {
        if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
            params.Page = p
        }
    }
    if sizeStr := c.QueryParam("limit"); sizeStr != "" {
        if s, err := strconv.Atoi(sizeStr); err == nil && s > 0 && s <= 100 {
            params.Size = s
        }
    }

    pageResult, err := h.probSvc.GetProblems(c.Request().Context(), filter, params)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusOK, pageResult)
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
    
    // In a real app we extract from JWT token context
    userIDStr := c.QueryParam("user_id")
    if userIDStr == "" {
        userIDStr = "00000000-0000-0000-0000-000000000001"
    }
    userID, err := uuid.Parse(userIDStr)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "invalid user id")
    }

    problem, err := h.probSvc.GetProblem(c.Request().Context(), slug)
    if err != nil {
        return echo.NewHTTPError(http.StatusNotFound, "problem not found")
    }

    isFavorited, err := h.probSvc.ToggleFavorite(c.Request().Context(), userID, problem.ID)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "failed to toggle favorite")
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "is_favorite": isFavorited,
    })
}
