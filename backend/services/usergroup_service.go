package services

import (
	"backend/db"
	"errors"

	"github.com/google/uuid"
)

type UserGroupService struct{}

func NewUserGroupService() *UserGroupService {
	return &UserGroupService{}
}

func (s *UserGroupService) JoinGroup(userID, groupID, role, state string) error {
	if userID == "" {
		return errors.New("user ID es requerido")
	}
	if groupID == "" {
		return errors.New("group ID es requerido")
	}

	if _, err := uuid.Parse(userID); err != nil {
		return errors.New("user ID invalido")
	}
	if _, err := uuid.Parse(groupID); err != nil {
		return errors.New("group ID invalido")
	}

	if role == "" {
		role = "USER"
	}
	if state == "" {
		state = "JOINED"
	}

	if role != "USER" && role != "ADMIN" {
		return errors.New("el rol debe ser USER o ADMIN")
	}

	if state != "JOINED" && state != "PENDING" && state != "BANNED" {
		return errors.New("el estado debe ser JOINED, PENDING, o BANNED")
	}

	id := uuid.New().String()

	err := db.JoinGroup(id, userID, groupID, role, state)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserGroupService) LeaveGroup(groupID, userID string) error {
	if userID == "" {
		return errors.New("user ID es requerido")
	}
	if groupID == "" {
		return errors.New("group ID es requerido")
	}

	if _, err := uuid.Parse(userID); err != nil {
		return errors.New("user ID invalido")
	}
	if _, err := uuid.Parse(groupID); err != nil {
		return errors.New("group ID invalido")
	}

	err := db.DeleteUserGroupEntry(groupID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserGroupService) ChangeUserRole(groupID, userID, role string) error {
	if userID == "" {
		return errors.New("user ID es requerido")
	}
	if groupID == "" {
		return errors.New("group ID es requerido")
	}
	if role == "" {
		return errors.New("role es requerido")
	}

	if _, err := uuid.Parse(userID); err != nil {
		return errors.New("user ID invalido")
	}
	if _, err := uuid.Parse(groupID); err != nil {
		return errors.New("group ID invalido")
	}

	if role != "USER" && role != "ADMIN" {
		return errors.New("el rol debe ser USER o ADMIN")
	}

	err := db.UpdateUserGroupRole(groupID, userID, role)
	if err != nil {
		return err
	}

	return nil
}
