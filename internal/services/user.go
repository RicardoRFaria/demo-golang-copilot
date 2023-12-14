package services

import (
	"demo-golang/internal/repository"
	"demo-golang/model"
	"fmt"
	"sort"
)

type UserService interface {
	GetUser(id int) (model.User, error)
	SaveUser(user model.User) error
	ListUsers() ([]model.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (u *userService) GetUser(id int) (model.User, error) {
	user, err := u.userRepository.GetUser(id)
	if err != nil {
		return model.User{}, fmt.Errorf("error getting user: %w", err)
	}
	if user.ID == 0 {
		return model.User{}, fmt.Errorf("user not found")
	}
	return user, nil
}

func (u *userService) SaveUser(user model.User) error {
	err := u.userRepository.SaveUser(user)
	if err != nil {
		return fmt.Errorf("error saving user: %w", err)
	}
	return nil
}

func (u *userService) ListUsers() ([]model.User, error) {
	users, err := u.userRepository.ListUsers()
	if err != nil {
		return nil, fmt.Errorf("error listing users: %w", err)
	}

	sort.Slice(users, func(i, j int) bool {
		return users[i].Age < users[j].Age
	})

	filteredUsers := make([]model.User, 0)
	for _, user := range users {
		if user.Age < 18 {
			filteredUsers = append(filteredUsers, user)
		}
	}

	return filteredUsers, nil
}
