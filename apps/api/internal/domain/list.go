package domain

import (
    "time"
    "github.com/google/uuid"
)

type ProblemList struct {
    ID          uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
    UserID      *uuid.UUID `json:"user_id,omitempty" gorm:"type:uuid"`
    Title       string    `json:"title"`
    Slug        string    `json:"slug" gorm:"unique;not null"`
    Description string    `json:"description,omitempty"`
    IsPublic    bool      `json:"is_public" gorm:"default:true"`
    IsCurated   bool      `json:"is_curated" gorm:"default:false"`
    CreatedAt   time.Time `json:"created_at" gorm:"default:now()"`
    UpdatedAt   time.Time `json:"updated_at" gorm:"default:now()"`

    // Relations
    Items []ProblemListItem `json:"items,omitempty" gorm:"foreignKey:ListID"`
    User  *User             `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

type ProblemListItem struct {
    ListID     uuid.UUID `json:"list_id" gorm:"type:uuid;primaryKey"`
    ProblemID  uuid.UUID `json:"problem_id" gorm:"type:uuid;primaryKey"`
    OrderIndex int32     `json:"order_index" gorm:"default:0"`
    AddedAt    time.Time `json:"added_at" gorm:"default:now()"`

    // Relation
    Problem *ProblemSummary `json:"problem,omitempty" gorm:"foreignKey:ProblemID"`
}

type Favorite struct {
    UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;primaryKey"`
    ProblemID uuid.UUID `json:"problem_id" gorm:"type:uuid;primaryKey"`
    CreatedAt time.Time `json:"created_at" gorm:"default:now()"`
}
