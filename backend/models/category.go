package models

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID            uuid.UUID  `json:"id" db:"id"`
	Name          string     `json:"name" db:"name"`
	Description   string     `json:"description" db:"description"`
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at" db:"deleted_at"`
	UserCreatedID uuid.UUID  `json:"user_created_id" db:"user_created_id"`
	GroupID       uuid.UUID  `json:"group_id" db:"group_id"`
}
