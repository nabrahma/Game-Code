package domain

import (
    "time"

    "github.com/google/uuid"
)

type UserRole string

const (
    UserRoleUser          UserRole = "user"
    UserRoleAdmin         UserRole = "admin"
    UserRoleContentEditor UserRole = "content_editor"
)

type User struct {
    ID            uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
    Email         string    `json:"email" gorm:"uniqueIndex;not null"`
    EmailVerified *time.Time `json:"email_verified,omitempty"`
    Name          *string   `json:"name,omitempty"`
    Username      string    `json:"username" gorm:"uniqueIndex;not null"`
    AvatarURL     *string   `json:"avatar_url,omitempty"`
    Role          UserRole  `json:"role" gorm:"type:user_role;default:'user';not null"`
    Bio           *string   `json:"bio,omitempty"`
    CreatedAt     time.Time `json:"created_at" gorm:"default:now();not null"`
    UpdatedAt     time.Time `json:"updated_at" gorm:"default:now();not null"`
}

func (User) TableName() string {
    return "users"
}

type UserProblemProgress struct {
    UserID       uuid.UUID  `json:"user_id" gorm:"type:uuid;primaryKey"`
    ProblemID    uuid.UUID  `json:"problem_id" gorm:"type:uuid;primaryKey"`
    Status       string     `json:"status" gorm:"default:'attempted';not null"`
    AttemptCount int32      `json:"attempt_count" gorm:"default:0;not null"`
    LastAttempt  time.Time  `json:"last_attempt" gorm:"default:now();not null"`
    SolvedAt     *time.Time `json:"solved_at,omitempty"`
}

func (UserProblemProgress) TableName() string {
    return "user_problem_progress"
}
