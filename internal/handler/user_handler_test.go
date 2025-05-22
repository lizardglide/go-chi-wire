package handler

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"go-chi-wire/internal/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

// mockUserService implements service.UserService
type mockUserService struct {
	mockGetUserByID func(id int) (model.User, error)
	mockGetAllUsers func() []model.User
}

func (m *mockUserService) GetUserByID(id int) (model.User, error) {
	return m.mockGetUserByID(id)
}

func (m *mockUserService) GetAllUsers() []model.User {
	return m.mockGetAllUsers()
}

func TestGetUser_Success(t *testing.T) {
	mockService := &mockUserService{
		mockGetUserByID: func(id int) (model.User, error) {
			return model.User{ID: id, Name: "TestUser"}, nil
		},
	}

	h := handler.NewUserHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/user/5", nil)
	rr := httptest.NewRecorder()

	// Use chi to mock the routing context
	router := chi.NewRouter()
	router.Get("/user/{id}", h.GetUser)
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", rr.Code)
	}

	var user model.User
	err := json.NewDecoder(rr.Body).Decode(&user)
	if err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if user.ID != 5 || user.Name != "TestUser" {
		t.Errorf("unexpected user: %+v", user)
	}
}

func TestGetUser_NotFound(t *testing.T) {
	mockService := &mockUserService{
		mockGetUserByID: func(id int) (model.User, error) {
			return model.User{}, errors.New("not found")
		},
	}

	h := handler.NewUserHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/user/99", nil)
	rr := httptest.NewRecorder()

	router := chi.NewRouter()
	router.Get("/user/{id}", h.GetUser)
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Fatalf("expected 404 Not Found, got %d", rr.Code)
	}
}

func TestGetAllUsers_Success(t *testing.T) {
	mockService := &mockUserService{
		mockGetAllUsers: func() []model.User {
			return []model.User{
				{ID: 1, Name: "Alice"},
				{ID: 2, Name: "Bob"},
			}
		},
	}

	h := handler.NewUserHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rr := httptest.NewRecorder()

	router := chi.NewRouter()
	router.Get("/users", h.GetAllUsers)
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", rr.Code)
	}

	var users []model.User
	if err := json.NewDecoder(rr.Body).Decode(&users); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(users) != 2 {
		t.Fatalf("expected 2 users, got %d", len(users))
	}
	if users[0].Name != "Alice" || users[1].Name != "Bob" {
		t.Errorf("unexpected users: %+v", users)
	}
}
