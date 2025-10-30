package models

import (
	"time"

	"github.com/google/uuid"
)

type Group struct {
	ID               uuid.UUID  `json:"id" db:"id"`
	Name             string     `json:"name" db:"name"`
	Description      string     `json:"description" db:"description"`
	UserCreatedID    uuid.UUID  `json:"user_created_id" db:"user_created_id"`
	CreationDate     time.Time  `json:"creation_date" db:"creation_date"`
	LastActivityDate *time.Time `json:"last_activity_date" db:"last_activity_date"`
	Privacy          string     `json:"privacy" db:"privacy"`
	State            string     `json:"state" db:"state"`
}
