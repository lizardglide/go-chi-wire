package service

import "go-chi-wire/internal/model"

type UserService interface {
	GetAllUsers() []model.User
	GetUserByID(id int) (model.User, error)
}
