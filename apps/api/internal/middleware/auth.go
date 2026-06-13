package middleware

import (
    "net/http"
    "strings"

    "github.com/labstack/echo/v4"
    "github.com/gc-platform/api/internal/config"
    "github.com/gc-platform/api/pkg/jwt"
)

const (
    AuthHeaderKey  = "Authorization"
    AuthTypeBearer = "Bearer"
    UserContextKey = "user_payload"
)

func RequireAuth(cfg *config.Config) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            tokenMaker, err := jwt.NewJWTMaker(cfg.JWTAccessSecret)
            if err != nil {
                return echo.NewHTTPError(http.StatusInternalServerError, "invalid token configuration")
            }

            // 1. Try cookie first
            cookie, err := c.Cookie("access_token")
            var tokenStr string
            if err == nil {
                tokenStr = cookie.Value
            } else {
                // 2. Try header
                authHeader := c.Request().Header.Get(AuthHeaderKey)
                if len(authHeader) == 0 {
                    return echo.NewHTTPError(http.StatusUnauthorized, "authorization header or cookie not provided")
                }
                fields := strings.Fields(authHeader)
                if len(fields) < 2 || strings.ToLower(fields[0]) != strings.ToLower(AuthTypeBearer) {
                    return echo.NewHTTPError(http.StatusUnauthorized, "invalid authorization format")
                }
                tokenStr = fields[1]
            }

            payload, err := tokenMaker.VerifyToken(tokenStr)
            if err != nil {
                return echo.NewHTTPError(http.StatusUnauthorized, "invalid or expired token")
            }

            c.Set(UserContextKey, payload)
            return next(c)
        }
    }
}

func RequireRole(role string) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            payload, ok := c.Get(UserContextKey).(*jwt.TokenPayload)
            if !ok {
                return echo.NewHTTPError(http.StatusUnauthorized, "user not authenticated")
            }

            if payload.Role != role && payload.Role != "admin" {
                return echo.NewHTTPError(http.StatusForbidden, "you don't have permission to access this resource")
            }

            return next(c)
        }
    }
}
