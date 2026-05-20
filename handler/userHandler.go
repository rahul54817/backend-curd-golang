package handler

import (
	"encoding/json"
	"go-backend/models"
	"go-backend/services"
	"net/http"
)

type UserHandler struct {
	Service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	id, err := h.Service.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User Created",
		"id":      id,
	})

}

func (h *UserHandler) GetUserList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	users, err := h.Service.GetUserList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")

	id := r.URL.Query().Get("id")

	user, err := h.Service.GetUserById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	id := r.URL.Query().Get("id")
	count, err := h.Service.DeleteUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if count == 0 {
		http.Error(w, "User not found", 404)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User deleted",
		"count":   count,
	})
}
