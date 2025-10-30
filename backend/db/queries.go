package db

import (
	"backend/models"

	"github.com/google/uuid"
)

func GetChannelsByCategory(groupID, categoryID uuid.UUID) ([]models.Channel, error) {
	var channels []models.Channel
	err := Database.Select(&channels, `SELECT * FROM "Channel" WHERE group_id = $1 AND category_id = $2`, groupID, categoryID)
	return channels, err
}

func GetAllUsersByGroup(groupID uuid.UUID) ([]models.UserGroup, error) {
	var users []models.UserGroup
	err := Database.Select(&users, `SELECT u.id, u.name, u.email, ug.role, ug.state FROM "User" u JOIN "UserGroup" ug ON u.id = ug.user_id WHERE ug.group_id = $1`, groupID)
	return users, err
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := Database.Select(&users, `SELECT id, name, email, profile_photo_url FROM "User"`)
	return users, err
}

func DeleteUserGroupEntry(groupID, userID string) error {
	_, err := Database.Exec(`DELETE FROM "UserGroup" WHERE group_id = $1 AND user_id = $2`, groupID, userID)
	return err
}

func UpdateUserGroupRole(groupID, userID, role string) error {
	_, err := Database.Exec(`UPDATE "UserGroup" SET role = $1 WHERE group_id = $2 AND user_id = $3`, role, groupID, userID)
	return err
}

func CreateGroup(id uuid.UUID, name, description, privacy, state string, userCreatedID uuid.UUID) error {
	_, err := Database.Exec(
		`INSERT INTO "Group" (id, name, description, privacy, state, user_created_id) VALUES ($1, $2, $3, $4, $5, $6)`,
		id, name, description, privacy, state, userCreatedID,
	)
	return err
}

func JoinGroup(id, userID, groupID, role, state string) error {
	_, err := Database.Exec(
		`INSERT INTO "UserGroup" (id, user_id, group_id, role, state) VALUES ($1, $2, $3, $4, $5)`,
		id, userID, groupID, role, state,
	)
	return err
}

func CreateCategory(id, name, description, userCreatedID, groupID string) error {
	_, err := Database.Exec(
		`INSERT INTO "Category" (id, name, description, user_created_id, group_id) VALUES ($1, $2, $3, $4, $5)`,
		id, name, description, userCreatedID, groupID,
	)
	return err
}

func CreateChannel(id, name, description, groupID, categoryID, channelState string) error {
	_, err := Database.Exec(
		`INSERT INTO "Channel" (id, name, description, group_id, category_id, channel_state) VALUES ($1, $2, $3, $4, $5, $6)`,
		id, name, description, groupID, categoryID, channelState,
	)
	return err
}

func CreateUser(id uuid.UUID, name, email, profilePhotoURL string) error {
	_, err := Database.Exec(
		`INSERT INTO "User" (id, name, email, profile_photo_url) VALUES ($1, $2, $3, $4)`,
		id, name, email, profilePhotoURL,
	)
	return err
}

func GetAllGroups() ([]models.Group, error) {
	var groups []models.Group
	err := Database.Select(&groups, `SELECT * FROM "Group"`)
	return groups, err
}

func GetCategoriesByGroup(groupID uuid.UUID) ([]models.Category, error) {
	var categories []models.Category
	err := Database.Select(&categories, `SELECT * FROM "Category" WHERE group_id = $1`, groupID)
	return categories, err
}

func GetChannelsByGroup(groupID uuid.UUID) ([]models.Channel, error) {
	var channels []models.Channel
	err := Database.Select(&channels, `SELECT * FROM "Channel" WHERE group_id = $1`, groupID)
	return channels, err
}

func DeleteGroup(id uuid.UUID) error {
	_, err := Database.Exec(`DELETE FROM "Group" WHERE id = $1`, id)
	return err
}

func DeleteCategory(id uuid.UUID) error {
	_, err := Database.Exec(`DELETE FROM "Category" WHERE id = $1`, id)
	return err
}

func DeleteChannel(id uuid.UUID) error {
	_, err := Database.Exec(`DELETE FROM "Channel" WHERE id = $1`, id)
	return err
}

func DeleteUser(id uuid.UUID) error {
	_, err := Database.Exec(`DELETE FROM "User" WHERE id = $1`, id)
	return err
}
