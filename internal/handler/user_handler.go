package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go-chi-wire/internal/service"
	"net/http"
	"strconv"
)

type UserHandler struct {
	Service service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{Service: userService}
}

func (userHandler *UserHandler) GetUserByID(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	userId := chi.URLParam(request, "id")
	id, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(responseWriter, "Invalid user id", http.StatusBadRequest)
		return
	}

	user, err := userHandler.Service.GetUserByID(id)
	if err != nil {
		http.Error(responseWriter, "User not found"+err.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(responseWriter).Encode(user); err != nil {
		http.Error(responseWriter, "Failed to encode response", http.StatusInternalServerError)
	}

}

func (userHandler *UserHandler) GetAllUsers(responseWriter http.ResponseWriter, request *http.Request) {
	users := userHandler.Service.GetAllUsers()
	if err := json.NewEncoder(responseWriter).Encode(users); err != nil {
		http.Error(responseWriter, "Failed to encode response", http.StatusInternalServerError)
	}
}
