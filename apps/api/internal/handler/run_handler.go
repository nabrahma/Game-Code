package handler

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/labstack/echo/v4"
    "github.com/gc-platform/api/internal/domain"
    "github.com/gc-platform/api/internal/service"
)

type RunHandler struct {
    runSvc service.RunService
}

func NewRunHandler(runSvc service.RunService) *RunHandler {
    return &RunHandler{runSvc: runSvc}
}

func (h *RunHandler) ExecuteCode(c echo.Context) error {
    var req domain.RunRequest
    if err := c.Bind(&req); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "invalid request payload")
    }

    if req.Code == "" {
        return echo.NewHTTPError(http.StatusBadRequest, "code is required")
    }

    runID, err := h.runSvc.EnqueueRun(c.Request().Context(), req)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "failed to enqueue execution task")
    }

    return c.JSON(http.StatusOK, map[string]string{
        "run_id": runID,
    })
}

func (h *RunHandler) StreamRunLogs(c echo.Context) error {
    runID := c.Param("runId")
    if runID == "" {
        return echo.NewHTTPError(http.StatusBadRequest, "runId is required")
    }

    // Set headers for SSE
    c.Response().Header().Set(echo.HeaderContentType, "text/event-stream")
    c.Response().Header().Set(echo.HeaderCacheControl, "no-cache")
    c.Response().Header().Set(echo.HeaderConnection, "keep-alive")

    ch, err := h.runSvc.SubscribeToRun(c.Request().Context(), runID)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "failed to subscribe to execution logs")
    }

    flusher, ok := c.Response().Writer.(http.Flusher)
    if !ok {
        return echo.NewHTTPError(http.StatusInternalServerError, "streaming unsupported")
    }

    for {
        select {
        case <-c.Request().Context().Done():
            return nil
        case event, ok := <-ch:
            if !ok {
                return nil
            }
            
            data, _ := json.Marshal(event)
            fmt.Fprintf(c.Response().Writer, "data: %s\n\n", data)
            flusher.Flush()
            
            if event.Status == string(domain.RunStatusSuccess) || 
               event.Status == string(domain.RunStatusError) || 
               event.Status == string(domain.RunStatusTimeout) {
                return nil
            }
        }
    }
}
