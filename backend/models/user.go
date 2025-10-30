package models

import "github.com/google/uuid"

type User struct {
	ID              uuid.UUID `db:"id" json:"id"`
	Name            string    `db:"name" json:"name"`
	Email           string    `db:"email" json:"email"`
	ProfilePhotoURL string    `db:"profile_photo_url" json:"profilePhotoURL"`
}
