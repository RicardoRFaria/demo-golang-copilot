package repository

import (
	"demo-golang/model"
	"fmt"
)

type UserRepository interface {
	GetUser(int) (model.User, error)
	SaveUser(user model.User) error
	ListUsers() ([]model.User, error)
}

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (u *userRepository) GetUser(int) (model.User, error) {
	return model.User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@gmail.com",
		Age:       25,
	}, nil
}

func (u *userRepository) SaveUser(user model.User) error {
	return fmt.Errorf("not implemented")
}

func (u *userRepository) ListUsers() ([]model.User, error) {
	panic("implement me")
}
