package middleware

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/gc-platform/api/internal/config"
    "github.com/redis/go-redis/v9"
    "go.uber.org/zap"
)

func Register(e *echo.Echo, cfg *config.Config, rdb *redis.Client, logger *zap.Logger) {
    e.Use(middleware.Recover())
    e.Use(middleware.RequestID())
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{cfg.FrontendURL, "http://localhost:3000"},
        AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, AuthHeaderKey},
        AllowCredentials: true,
    }))
    
    // Minimal zap logger middleware
    e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            err := next(c)
            if err != nil {
                c.Error(err)
            }
            logger.Info("request",
                zap.String("method", c.Request().Method),
                zap.String("uri", c.Request().RequestURI),
                zap.Int("status", c.Response().Status),
            )
            return nil
        }
    })
}
