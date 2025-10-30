package handlers

import (
	"backend/services"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

var (
	groupService     *services.GroupService
	userService      *services.UserService
	userGroupService *services.UserGroupService
)

func InitializeServices() {
	groupService = services.NewGroupService()
	userService = services.NewUserService()
	userGroupService = services.NewUserGroupService()
}

func GetChannelByCategory(w http.ResponseWriter, r *http.Request) {
	groupIDStr := chi.URLParam(r, "groupID")
	categoryIDStr := chi.URLParam(r, "categoryID")
	groupID, err := uuid.Parse(groupIDStr)
	if err != nil {
		http.Error(w, "group ID invalido", http.StatusBadRequest)
		return
	}
	categoryID, err := uuid.Parse(categoryIDStr)
	if err != nil {
		http.Error(w, "category ID invalida", http.StatusBadRequest)
		return
	}
	channels, err := groupService.GetChannelsByCategory(groupID, categoryID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(channels)
}

func GetAllUsersByGroup(w http.ResponseWriter, r *http.Request) {
	groupIDStr := chi.URLParam(r, "groupID")
	groupID, err := uuid.Parse(groupIDStr)
	if err != nil {
		http.Error(w, "group ID invalido", http.StatusBadRequest)
		return
	}
	users, err := userService.GetAllUsersByGroup(groupID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetGroups(w http.ResponseWriter, r *http.Request) {
	groups, err := groupService.GetAllGroups()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(groups)
}

func CreateGroup(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Privacy     string `json:"privacy"`
		State       string `json:"state"`
		UserID      string `json:"userID"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "request payload invalido", http.StatusBadRequest)
		return
	}

	userCreatedID, err := uuid.Parse(req.UserID)
	if err != nil {
		http.Error(w, "user ID invalido", http.StatusBadRequest)
		return
	}

	groupID, err := groupService.CreateGroup(req.Name, req.Description, req.Privacy, req.State, userCreatedID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"groupID": groupID.String(),
	})
}

func JoinGroup(w http.ResponseWriter, r *http.Request) {
	groupID := chi.URLParam(r, "groupID")
	var user struct {
		UserID string `json:"userID"`
		Role   string `json:"role"`
		State  string `json:"state"`
	}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "request payload invalido", http.StatusBadRequest)
		return
	}

	err := userGroupService.JoinGroup(user.UserID, groupID, user.Role, user.State)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	groupID := chi.URLParam(r, "groupID")
	var category struct {
		Name          string `json:"name"`
		Description   string `json:"description"`
		UserCreatedID string `json:"userCreatedID"`
	}
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, "request payload invalido", http.StatusBadRequest)
		return
	}

	categoryID, err := groupService.CreateCategory(category.Name, category.Description, category.UserCreatedID, groupID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"id": categoryID,
	})
}

func CreateChannel(w http.ResponseWriter, r *http.Request) {
	groupID := chi.URLParam(r, "groupID")
	categoryID := chi.URLParam(r, "categoryID")
	var channel struct {
		Name         string `json:"name"`
		Description  string `json:"description"`
		ChannelState string `json:"channelState"`
		UserID       string `json:"userID"`
	}
	if err := json.NewDecoder(r.Body).Decode(&channel); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	channelID, err := groupService.CreateChannel(channel.Name, channel.Description, groupID, categoryID, channel.ChannelState)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"id": channelID,
	})
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name            string `json:"name"`
		Email           string `json:"email"`
		ProfilePhotoURL string `json:"profilePhotoURL"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "request payload invalido", http.StatusBadRequest)
		return
	}

	userID, err := userService.CreateUser(req.Name, req.Email, req.ProfilePhotoURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"userID": userID.String(),
	})
}

func GetCategoriesByGroup(w http.ResponseWriter, r *http.Request) {
	groupIDStr := chi.URLParam(r, "groupID")
	groupID, err := uuid.Parse(groupIDStr)
	if err != nil {
		http.Error(w, "Id de grupo invalido", http.StatusBadRequest)
		return
	}

	categories, err := groupService.GetCategoriesByGroup(groupID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func GetChannelsByGroup(w http.ResponseWriter, r *http.Request) {
	groupIDStr := chi.URLParam(r, "groupID")
	groupID, err := uuid.Parse(groupIDStr)
	if err != nil {
		http.Error(w, "Id de grupo invalido", http.StatusBadRequest)
		return
	}

	channels, err := groupService.GetChannelsByGroup(groupID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(channels)
}

func DeleteGroup(w http.ResponseWriter, r *http.Request) {
	groupIDStr := chi.URLParam(r, "groupID")
	groupID, err := uuid.Parse(groupIDStr)
	if err != nil {
		http.Error(w, "group ID invalido", http.StatusBadRequest)
		return
	}

	if err := groupService.DeleteGroup(groupID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	categoryIDStr := chi.URLParam(r, "categoryID")
	categoryID, err := uuid.Parse(categoryIDStr)
	if err != nil {
		http.Error(w, "category id invalido", http.StatusBadRequest)
		return
	}

	if err := groupService.DeleteCategory(categoryID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteChannel(w http.ResponseWriter, r *http.Request) {
	channelIDStr := chi.URLParam(r, "channelID")
	channelID, err := uuid.Parse(channelIDStr)
	if err != nil {
		http.Error(w, "channel id invalido", http.StatusBadRequest)
		return
	}

	if err := groupService.DeleteChannel(channelID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "userID")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(w, "User ID invalido", http.StatusBadRequest)
		return
	}

	if err := userService.DeleteUser(userID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
