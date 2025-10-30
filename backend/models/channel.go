package models

import (
	"time"

	"github.com/google/uuid"
)

type Channel struct {
	ID           uuid.UUID `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	Description  string    `json:"description" db:"description"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	GroupID      uuid.UUID `json:"group_id" db:"group_id"`
	CategoryID   uuid.UUID `json:"category_id" db:"category_id"`
	ChannelState string    `json:"channel_state" db:"channel_state"`
}
