package repository

import "belajar-go-echo/app/model"

// UserRepository interface untuk operasi user
type UserRepository interface {
	GetAllUsers() ([]model.User, error)
	CreateUser(user *model.User) (*model.User, error)

	// inteface untuk get all users with auth
	GetAuthenticatedUsers() ([]model.User, error)
}
