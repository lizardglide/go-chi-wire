package util

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"go-chi-wire/internal/handler"
	"net/http"
)

func NewUserRouter(handler *handler.UserHandler) http.Handler {
	r := chi.NewRouter()

	// CORS middleware
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}, // Adjust in prod
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // Max age in seconds
	}))

	// Global middleware to set content-type
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	r.Get("/user/{id}", handler.GetUser)
	r.Get("/users", handler.GetAllUsers)

	return r
}

//package service
//
//import (
//"testing"
//"go-chi-wire/internal/model"
//)
//
//func TestGetUserByID(t *testing.T) {
//svc := NewUserService()
//
//user, err := svc.GetUserByID(1)
//if err != nil {
//t.Fatalf("Expected user, got error: %v", err)
//}
//
//if user.ID != 1 || user.Name != "Alice" {
//t.Errorf("Unexpected user returned: %+v", user)
//}
//}
//
//func TestGetUserByID_NotFound(t *testing.T) {
//svc := NewUserService()
//
//_, err := svc.GetUserByID(999)
//if err == nil {
//t.Fatal("Expected error, got nil")
//}
//}
//
//func TestGetAllUsers(t *testing.T) {
//svc := NewUserService()
//
//users := svc.GetAllUsers()
//if len(users) != 2 {
//t.Errorf("Expected 2 users, got %d", len(users))
//}
//}
//
//Abstracted response
//
//package util
//
//import (
//"encoding/json"
//"net/http"
//)
//
//type ErrorResponse struct {
//Error string `json:"error"`
//}
//
//func WriteJSON(w http.ResponseWriter, statusCode int, data interface{}) {
//w.Header().Set("Content-Type", "application/json")
//w.WriteHeader(statusCode)
//json.NewEncoder(w).Encode(data)
//}
//
//func WriteError(w http.ResponseWriter, statusCode int, msg string) {
//WriteJSON(w, statusCode, ErrorResponse{Error: msg})
//}
//
