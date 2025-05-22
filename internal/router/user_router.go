package router

import (
	"github.com/go-chi/chi/v5"
	"go-chi-wire/internal/handler"
	"net/http"
)

func NewUserRouter(handler *handler.UserHandler) http.Handler {
	router := chi.NewRouter()
	router.Get("/user/{id}", handler.GetUserByID)
	router.Get("/users", handler.GetAllUsers)
	return router
}
