package service

import (
	"testing"
)

func TestGetUserByID(t *testing.T) {
	userService := NewUserService()

	user, err := userService.GetUserByID(1)
	if err != nil {
		t.Fatalf("TestGetUserByID err : %v", err)
	}

	if user.ID != 1 || user.Name != "Alice" {
		t.Errorf("Unexpected user returned : %v", user)
	}
}
