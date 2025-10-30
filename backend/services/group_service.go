package services

import (
	"backend/db"
	"backend/models"
	"errors"

	"github.com/google/uuid"
)

type GroupService struct{}

func NewGroupService() *GroupService {
	return &GroupService{}
}

func (s *GroupService) CreateGroup(name, description, privacy, state string, userCreatedID uuid.UUID) (uuid.UUID, error) {
	if name == "" {
		return uuid.Nil, errors.New("group name es requerido")
	}
	if userCreatedID == uuid.Nil {
		return uuid.Nil, errors.New("user creator ID es requerido")
	}

	if privacy == "" {
		privacy = "PUBLIC"
	}
	if state == "" {
		state = "ACTIVE"
	}

	if privacy != "PUBLIC" && privacy != "PRIVATE" {
		return uuid.Nil, errors.New("la privacidad puede ser PUBLIC o PRIVATE")
	}

	if state != "ACTIVE" && state != "INACTIVE" && state != "ARCHIVED" {
		return uuid.Nil, errors.New("el estado debe ser ACTIVE, INACTIVE, o ARCHIVED")
	}

	groupID := uuid.New()

	err := db.CreateGroup(groupID, name, description, privacy, state, userCreatedID)
	if err != nil {
		return uuid.Nil, err
	}

	return groupID, nil
}

func (s *GroupService) GetAllGroups() ([]models.Group, error) {
	groups, err := db.GetAllGroups()
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func (s *GroupService) DeleteGroup(groupID uuid.UUID) error {
	if groupID == uuid.Nil {
		return errors.New("group ID es requerido")
	}

	err := db.DeleteGroup(groupID)
	if err != nil {
		return err
	}

	return nil
}

func (s *GroupService) CreateCategory(name, description, userCreatedID, groupID string) (string, error) {
	if name == "" {
		return "", errors.New("el nombre de la categoria es requerido")
	}
	if groupID == "" {
		return "", errors.New("group ID es requerido")
	}
	if userCreatedID == "" {
		return "", errors.New("user created ID es requerido")
	}

	if _, err := uuid.Parse(groupID); err != nil {
		return "", errors.New("group ID invalido")
	}
	if _, err := uuid.Parse(userCreatedID); err != nil {
		return "", errors.New("user creator ID invalido")
	}

	categoryID := uuid.New().String()

	err := db.CreateCategory(categoryID, name, description, userCreatedID, groupID)
	if err != nil {
		return "", err
	}

	return categoryID, nil
}

func (s *GroupService) GetCategoriesByGroup(groupID uuid.UUID) ([]models.Category, error) {
	if groupID == uuid.Nil {
		return nil, errors.New("group ID es requerido")
	}

	categories, err := db.GetCategoriesByGroup(groupID)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (s *GroupService) DeleteCategory(categoryID uuid.UUID) error {
	if categoryID == uuid.Nil {
		return errors.New("category ID es requerido")
	}

	err := db.DeleteCategory(categoryID)
	if err != nil {
		return err
	}

	return nil
}

func (s *GroupService) CreateChannel(name, description, groupID, categoryID, channelState string) (string, error) {
	if name == "" {
		return "", errors.New("el nombre del canal es requerido")
	}
	if groupID == "" {
		return "", errors.New("group ID es requerido")
	}
	if categoryID == "" {
		return "", errors.New("category ID es requerido")
	}

	if _, err := uuid.Parse(groupID); err != nil {
		return "", errors.New("group ID invalido")
	}
	if _, err := uuid.Parse(categoryID); err != nil {
		return "", errors.New("category ID invalido")
	}

	if channelState == "" {
		channelState = "ACTIVE"
	}

	if channelState != "ACTIVE" && channelState != "INACTIVE" && channelState != "ARCHIVED" {
		return "", errors.New("el estado de un canal debe ser ACTIVE, INACTIVE, o ARCHIVED")
	}

	channelID := uuid.New().String()

	err := db.CreateChannel(channelID, name, description, groupID, categoryID, channelState)
	if err != nil {
		return "", err
	}

	return channelID, nil
}

func (s *GroupService) GetChannelsByGroup(groupID uuid.UUID) ([]models.Channel, error) {
	if groupID == uuid.Nil {
		return nil, errors.New("group ID es requerido")
	}

	channels, err := db.GetChannelsByGroup(groupID)
	if err != nil {
		return nil, err
	}

	return channels, nil
}

func (s *GroupService) GetChannelsByCategory(groupID, categoryID uuid.UUID) ([]models.Channel, error) {
	if groupID == uuid.Nil {
		return nil, errors.New("group ID es requerido")
	}
	if categoryID == uuid.Nil {
		return nil, errors.New("category ID es requerido")
	}

	channels, err := db.GetChannelsByCategory(groupID, categoryID)
	if err != nil {
		return nil, err
	}

	return channels, nil
}

func (s *GroupService) DeleteChannel(channelID uuid.UUID) error {
	if channelID == uuid.Nil {
		return errors.New("channel ID es requerido")
	}

	err := db.DeleteChannel(channelID)
	if err != nil {
		return err
	}

	return nil
}
