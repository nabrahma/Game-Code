package domain

import (
    "time"
    "github.com/google/uuid"
)

type DiscussPost struct {
    ID        uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
    ProblemID uuid.UUID `json:"problem_id" gorm:"type:uuid;not null;index"`
    UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null;index"`
    Title     string    `json:"title" gorm:"not null"`
    Content   string    `json:"content" gorm:"not null"`
    Upvotes   int32     `json:"upvotes" gorm:"default:0"`
    CreatedAt time.Time `json:"created_at" gorm:"default:now()"`
    UpdatedAt time.Time `json:"updated_at" gorm:"default:now()"`

    // Relations
    User     *User            `json:"user,omitempty" gorm:"foreignKey:UserID"`
    Comments []DiscussComment `json:"comments,omitempty" gorm:"foreignKey:PostID"`
}

type DiscussComment struct {
    ID        uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
    PostID    uuid.UUID `json:"post_id" gorm:"type:uuid;not null;index"`
    UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
    Content   string    `json:"content" gorm:"not null"`
    Upvotes   int32     `json:"upvotes" gorm:"default:0"`
    CreatedAt time.Time `json:"created_at" gorm:"default:now()"`
    UpdatedAt time.Time `json:"updated_at" gorm:"default:now()"`

    // Relations
    User *User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

type PostUpvote struct {
    PostID    uuid.UUID `json:"post_id" gorm:"type:uuid;primaryKey"`
    UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;primaryKey"`
    CreatedAt time.Time `json:"created_at" gorm:"default:now()"`
}

type CommentUpvote struct {
    CommentID uuid.UUID `json:"comment_id" gorm:"type:uuid;primaryKey"`
    UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;primaryKey"`
    CreatedAt time.Time `json:"created_at" gorm:"default:now()"`
}
