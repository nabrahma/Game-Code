package middleware

import (
    "fmt"
    "net/http"

    "github.com/getsentry/sentry-go"
    "github.com/labstack/echo/v4"
)

// SentryMiddleware captures panics and routes them to Sentry.
func SentryMiddleware() echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            defer func() {
                if err := recover(); err != nil {
                    eventID := sentry.CurrentHub().Recover(err)
                    if eventID != nil {
                        sentry.Flush(2 * 1000 * 1000 * 1000) // 2 seconds
                    }
                    // Re-panic after capturing to let other middlewares handle it (like default Recover)
                    c.Error(fmt.Errorf("panic: %v", err))
                }
            }()
            
            err := next(c)
            if err != nil {
                // Ignore 400-level errors, capture 500-level errors
                if he, ok := err.(*echo.HTTPError); ok {
                    if he.Code >= 500 {
                        sentry.CaptureException(err)
                    }
                } else {
                    sentry.CaptureException(err)
                }
            }
            return err
        }
    }
}
