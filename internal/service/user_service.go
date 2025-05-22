package service

import (
	"fmt"
	"go-chi-wire/internal/model"
)

var users = []model.User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Peach"},
}

type userServiceImpl struct {
}

func NewUserService() UserService {
	return &userServiceImpl{}
}

func (userServiceImpl userServiceImpl) GetUserByID(id int) (model.User, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}
	return model.User{}, fmt.Errorf("User not found")
}

func (userServiceImpl userServiceImpl) GetAllUsers() []model.User {
	return users
}
