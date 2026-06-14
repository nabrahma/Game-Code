package handler

import (
    "net/http"

    "github.com/google/uuid"
    "github.com/labstack/echo/v4"
    "github.com/gc-platform/api/internal/domain"
    "github.com/gc-platform/api/internal/service"
)

type AdminHandler struct {
    probService service.ProblemService
}

func NewAdminHandler(probService service.ProblemService) *AdminHandler {
    return &AdminHandler{probService: probService}
}

func (h *AdminHandler) CreateProblem(c echo.Context) error {
    var p domain.Problem
    if err := c.Bind(&p); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
    }
    
    // Auto-generate UUID if not provided
    if p.ID == uuid.Nil {
        p.ID = uuid.New()
    }

    if err := h.probService.CreateProblem(c.Request().Context(), &p); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusCreated, p)
}

func (h *AdminHandler) UpdateProblem(c echo.Context) error {
    slug := c.Param("slug")
    var updates map[string]interface{}
    if err := c.Bind(&updates); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
    }

    if err := h.probService.UpdateProblem(c.Request().Context(), slug, updates); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, map[string]string{"status": "updated"})
}

func (h *AdminHandler) DeleteProblem(c echo.Context) error {
    idStr := c.Param("id")
    id, err := uuid.Parse(idStr)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid uuid"})
    }

    if err := h.probService.DeleteProblem(c.Request().Context(), id); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func (h *AdminHandler) UpsertTestCase(c echo.Context) error {
    var tc domain.TestCase
    if err := c.Bind(&tc); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
    }
    
    if err := h.probService.UpsertTestCase(c.Request().Context(), &tc); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, tc)
}

func (h *AdminHandler) DeleteTestCase(c echo.Context) error {
    idStr := c.Param("id")
    id, err := uuid.Parse(idStr)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid uuid"})
    }

    if err := h.probService.DeleteTestCase(c.Request().Context(), id); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func (h *AdminHandler) UpsertStarterCode(c echo.Context) error {
    var sc domain.StarterCode
    if err := c.Bind(&sc); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
    }

    if err := h.probService.UpsertStarterCode(c.Request().Context(), &sc); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, sc)
}

func (h *AdminHandler) UpsertEditorial(c echo.Context) error {
    var e domain.Editorial
    if err := c.Bind(&e); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
    }

    if err := h.probService.UpsertEditorial(c.Request().Context(), &e); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, e)
}
