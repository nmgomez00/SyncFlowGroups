package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := userService.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func LeftGroup(w http.ResponseWriter, r *http.Request) {
	groupID := chi.URLParam(r, "groupID")
	userID := chi.URLParam(r, "userID")

	err := userGroupService.LeaveGroup(groupID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func ChangeRole(w http.ResponseWriter, r *http.Request) {
	groupID := chi.URLParam(r, "groupID")
	userID := chi.URLParam(r, "userID")
	var req struct {
		Role string `json:"role"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "request payload invalido", http.StatusBadRequest)
		return
	}

	err := userGroupService.ChangeUserRole(groupID, userID, req.Role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
