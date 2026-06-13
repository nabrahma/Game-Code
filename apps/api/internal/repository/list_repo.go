package repository

import (
    "gorm.io/gorm"
)

type ListRepo interface {}

type listRepo struct {
    db *gorm.DB
}

func NewListRepo(db *gorm.DB) ListRepo {
    return &listRepo{db: db}
}
