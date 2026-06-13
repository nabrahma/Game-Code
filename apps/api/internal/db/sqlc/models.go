package sqlc

import (
    "context"
    "time"

    "github.com/google/uuid"
    "github.com/jackc/pgx/v5"
    "github.com/jackc/pgx/v5/pgtype"
)

type User struct {
    ID            uuid.UUID  `json:"id"`
    Email         string     `json:"email"`
    EmailVerified *time.Time `json:"email_verified"`
    Name          *string    `json:"name"`
    Username      string     `json:"username"`
    AvatarUrl     *string    `json:"avatar_url"`
    Role          string     `json:"role"`
    Bio           *string    `json:"bio"`
    CreatedAt     time.Time  `json:"created_at"`
    UpdatedAt     time.Time  `json:"updated_at"`
}

type OauthAccount struct {
    ID                uuid.UUID  `json:"id"`
    UserID            uuid.UUID  `json:"user_id"`
    Provider          string     `json:"provider"`
    ProviderAccountID string     `json:"provider_account_id"`
    AccessToken       *string    `json:"access_token"`
    RefreshToken      *string    `json:"refresh_token"`
    ExpiresAt         *time.Time `json:"expires_at"`
    CreatedAt         time.Time  `json:"created_at"`
}

type RefreshToken struct {
    ID        uuid.UUID  `json:"id"`
    UserID    uuid.UUID  `json:"user_id"`
    TokenHash string     `json:"token_hash"`
    ExpiresAt time.Time  `json:"expires_at"`
    RevokedAt *time.Time `json:"revoked_at"`
    CreatedAt time.Time  `json:"created_at"`
}

type DBTX interface {
    Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
    Query(context.Context, string, ...interface{}) (pgx.Rows, error)
    QueryRow(context.Context, string, ...interface{}) pgx.Row
}

func New(db DBTX) *Queries {
    return &Queries{db: db}
}

type Queries struct {
    db DBTX
}
