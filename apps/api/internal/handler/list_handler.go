package handler

import (
    "net/http"

    "github.com/google/uuid"
    "github.com/labstack/echo/v4"
    "github.com/gc-platform/api/internal/domain"
    "github.com/gc-platform/api/internal/service"
)

type ListHandler struct {
    listSvc service.ListService
}

func NewListHandler(listSvc service.ListService) *ListHandler {
    return &ListHandler{listSvc: listSvc}
}

func (h *ListHandler) GetCuratedLists(c echo.Context) error {
    lists, err := h.listSvc.GetCuratedLists(c.Request().Context())
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get curated lists")
    }
    return c.JSON(http.StatusOK, lists)
}

func (h *ListHandler) GetUserLists(c echo.Context) error {
    // In a real app, we extract userID from the JWT token
    // For now, let's assume we pass it via query or stub it
    userIDStr := c.QueryParam("user_id")
    if userIDStr == "" {
        // Stub user ID for now if not provided
        userIDStr = "00000000-0000-0000-0000-000000000001"
    }

    userID, err := uuid.Parse(userIDStr)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
    }

    lists, err := h.listSvc.GetUserLists(c.Request().Context(), userID)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get user lists")
    }
    return c.JSON(http.StatusOK, lists)
}

func (h *ListHandler) GetBySlug(c echo.Context) error {
    slug := c.Param("slug")
    if slug == "" {
        return echo.NewHTTPError(http.StatusBadRequest, "Slug is required")
    }

    list, err := h.listSvc.GetListBySlug(c.Request().Context(), slug)
    if err != nil {
        return echo.NewHTTPError(http.StatusNotFound, "List not found")
    }
    return c.JSON(http.StatusOK, list)
}

func (h *ListHandler) CreateList(c echo.Context) error {
    var req struct {
        Title       string `json:"title"`
        Description string `json:"description"`
        IsPublic    bool   `json:"is_public"`
    }
    if err := c.Bind(&req); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
    }

    // Stub user ID
    userID := uuid.MustParse("00000000-0000-0000-0000-000000000001")

    // Generate simple slug from title
    // Real implementation would use a proper slugifier
    slug := req.Title // Placeholder

    list := &domain.ProblemList{
        UserID:      &userID,
        Title:       req.Title,
        Slug:        slug,
        Description: req.Description,
        IsPublic:    req.IsPublic,
        IsCurated:   false,
    }

    if err := h.listSvc.CreateList(c.Request().Context(), list); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create list")
    }

    return c.JSON(http.StatusCreated, list)
}

func (h *ListHandler) AddProblem(c echo.Context) error {
    listIDStr := c.Param("id")
    listID, err := uuid.Parse(listIDStr)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid list ID")
    }

    var req struct {
        ProblemID string `json:"problem_id"`
    }
    if err := c.Bind(&req); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
    }
    problemID, err := uuid.Parse(req.ProblemID)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid problem ID")
    }

    if err := h.listSvc.AddProblemToList(c.Request().Context(), listID, problemID); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to add problem to list")
    }

    return c.NoContent(http.StatusOK)
}

func (h *ListHandler) RemoveProblem(c echo.Context) error {
    listIDStr := c.Param("id")
    listID, err := uuid.Parse(listIDStr)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid list ID")
    }

    problemIDStr := c.Param("problemId")
    problemID, err := uuid.Parse(problemIDStr)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid problem ID")
    }

    if err := h.listSvc.RemoveProblemFromList(c.Request().Context(), listID, problemID); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to remove problem from list")
    }

    return c.NoContent(http.StatusOK)
}
