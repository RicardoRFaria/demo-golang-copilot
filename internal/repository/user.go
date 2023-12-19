package repository

import (
	"context"
	"demo-golang/model"
	"fmt"
)

type UserRepository interface {
	GetUser(context.Context, int) (model.User, error)
	SaveUser(ctx context.Context, user model.User) error
	ListUsers(context.Context) ([]model.User, error)
}

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (u *userRepository) GetUser(context.Context, int) (model.User, error) {
	return model.User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@gmail.com",
		Age:       25,
	}, nil
}

func (u *userRepository) SaveUser(ctx context.Context, user model.User) error {
	return fmt.Errorf("not implemented")
}

func (u *userRepository) ListUsers(context.Context) ([]model.User, error) {
	panic("implement me")
}
