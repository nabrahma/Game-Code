package handler

import (
    "net/http"

    "github.com/google/uuid"
    "github.com/labstack/echo/v4"
    "github.com/gc-platform/api/internal/service"
)

type DiscussHandler struct {
    discussService service.DiscussService
}

func NewDiscussHandler(discussService service.DiscussService) *DiscussHandler {
    return &DiscussHandler{discussService: discussService}
}

func (h *DiscussHandler) ListPostsByProblem(c echo.Context) error {
    slug := c.Param("slug")
    posts, err := h.discussService.ListPostsByProblem(c.Request().Context(), slug)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, posts)
}

func (h *DiscussHandler) GetPostByID(c echo.Context) error {
    idStr := c.Param("id")
    postID, err := uuid.Parse(idStr)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid post id"})
    }

    post, err := h.discussService.GetPostByID(c.Request().Context(), postID)
    if err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "post not found"})
    }
    return c.JSON(http.StatusOK, post)
}

func (h *DiscussHandler) CreatePost(c echo.Context) error {
    slug := c.Param("slug")
    var req struct {
        UserID  uuid.UUID `json:"user_id"`
        Title   string    `json:"title"`
        Content string    `json:"content"`
    }
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
    }
    if req.UserID == uuid.Nil {
        // Fallback for demo
        req.UserID = uuid.MustParse("00000000-0000-0000-0000-000000000001")
    }

    post, err := h.discussService.CreatePost(c.Request().Context(), slug, req.UserID, req.Title, req.Content)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusCreated, post)
}

func (h *DiscussHandler) CreateComment(c echo.Context) error {
    idStr := c.Param("id")
    postID, err := uuid.Parse(idStr)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid post id"})
    }

    var req struct {
        UserID  uuid.UUID `json:"user_id"`
        Content string    `json:"content"`
    }
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
    }
    if req.UserID == uuid.Nil {
        // Fallback for demo
        req.UserID = uuid.MustParse("00000000-0000-0000-0000-000000000001")
    }

    comment, err := h.discussService.CreateComment(c.Request().Context(), postID, req.UserID, req.Content)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusCreated, comment)
}

func (h *DiscussHandler) ListCommentsForPost(c echo.Context) error {
    idStr := c.Param("id")
    postID, err := uuid.Parse(idStr)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid post id"})
    }

    comments, err := h.discussService.ListCommentsForPost(c.Request().Context(), postID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, comments)
}

func (h *DiscussHandler) TogglePostUpvote(c echo.Context) error {
    idStr := c.Param("id")
    postID, err := uuid.Parse(idStr)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid post id"})
    }

    userID := uuid.MustParse("00000000-0000-0000-0000-000000000001") // TODO: get from context
    added, err := h.discussService.TogglePostUpvote(c.Request().Context(), postID, userID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, map[string]bool{"added": added})
}

func (h *DiscussHandler) ToggleCommentUpvote(c echo.Context) error {
    idStr := c.Param("commentId")
    commentID, err := uuid.Parse(idStr)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid comment id"})
    }

    userID := uuid.MustParse("00000000-0000-0000-0000-000000000001") // TODO: get from context
    added, err := h.discussService.ToggleCommentUpvote(c.Request().Context(), commentID, userID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, map[string]bool{"added": added})
}
