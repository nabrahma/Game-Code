package handler

import (
    "net/http"

    "github.com/google/uuid"
    "github.com/labstack/echo/v4"
    "github.com/gc-platform/api/internal/service"
)

type SubmissionHandler struct {
    subSvc service.SubmissionService
}

func NewSubmissionHandler(subSvc service.SubmissionService) *SubmissionHandler {
    return &SubmissionHandler{subSvc: subSvc}
}

type CreateSubmissionRequest struct {
    ProblemID string `json:"problem_id" validate:"required"`
    Language  string `json:"language" validate:"required"`
    Code      string `json:"code" validate:"required"`
}

func (h *SubmissionHandler) Create(c echo.Context) error {
    var req CreateSubmissionRequest
    if err := c.Bind(&req); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    probID, err := uuid.Parse(req.ProblemID)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "invalid problem_id")
    }

    // Mock UserID until auth is implemented in middleware
    userID := uuid.MustParse("00000000-0000-0000-0000-000000000001")
    userIDStr := c.Request().Header.Get("X-User-ID")
    if userIDStr != "" {
        if parsed, err := uuid.Parse(userIDStr); err == nil {
            userID = parsed
        }
    }

    sub, err := h.subSvc.Create(c.Request().Context(), userID, probID, req.Language, req.Code)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusAccepted, sub)
}

func (h *SubmissionHandler) GetByID(c echo.Context) error {
    idStr := c.Param("id")
    id, err := uuid.Parse(idStr)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "invalid submission id")
    }

    sub, err := h.subSvc.GetByID(c.Request().Context(), id)
    if err != nil {
        return echo.NewHTTPError(http.StatusNotFound, "submission not found")
    }

    return c.JSON(http.StatusOK, sub)
}

func (h *SubmissionHandler) List(c echo.Context) error {
    return c.JSON(http.StatusOK, map[string]interface{}{"data": []interface{}{}})
}
