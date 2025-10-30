package services

import (
	"backend/db"
	"backend/models"
	"errors"

	"github.com/google/uuid"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) CreateUser(name, email, profilePhotoURL string) (uuid.UUID, error) {
	if name == "" {
		return uuid.Nil, errors.New("el nombre de usuario es requerido")
	}
	if email == "" {
		return uuid.Nil, errors.New("el email es requerido")
	}

	userID := uuid.New()

	err := db.CreateUser(userID, name, email, profilePhotoURL)
	if err != nil {
		return uuid.Nil, err
	}

	return userID, nil
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	users, err := db.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) DeleteUser(userID uuid.UUID) error {
	if userID == uuid.Nil {
		return errors.New("user ID es requerido")
	}

	err := db.DeleteUser(userID)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetAllUsersByGroup(groupID uuid.UUID) ([]models.UserGroup, error) {
	if groupID == uuid.Nil {
		return nil, errors.New("group ID es requerido")
	}

	users, err := db.GetAllUsersByGroup(groupID)
	if err != nil {
		return nil, err
	}

	return users, nil
}
