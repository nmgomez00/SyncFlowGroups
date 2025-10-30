package models

type UserGroup struct {
	ID    string `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
	Role  string `db:"role" json:"role"`
	State string `db:"state" json:"state"`
}
