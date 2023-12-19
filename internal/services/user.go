package services

import (
	"context"
	"demo-golang/internal/repository"
	"demo-golang/model"
	"fmt"
	"sort"
	"strings"
)

type UserService interface {
	GetUser(ctx context.Context, id int) (model.User, error)
	SaveUser(ctx context.Context, user model.User) error
	ListUsers(context.Context) ([]model.User, error)
	ListUserLastNames(context.Context) ([]string, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (u *userService) GetUser(ctx context.Context, id int) (model.User, error) {
	user, err := u.userRepository.GetUser(ctx, id)
	if err != nil {
		return model.User{}, fmt.Errorf("error getting user: %w", err)
	}
	if user.ID == 0 {
		return model.User{}, fmt.Errorf("user not found")
	}
	return user, nil
}

func (u *userService) SaveUser(ctx context.Context, user model.User) error {
	err := u.userRepository.SaveUser(ctx, user)
	if err != nil {
		return fmt.Errorf("error saving user: %w", err)
	}
	return nil
}

func (u *userService) ListUsers(ctx context.Context) ([]model.User, error) {
	users, err := u.userRepository.ListUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("error listing users: %w", err)
	}

	sort.Slice(users, func(i, j int) bool {
		return users[i].Age < users[j].Age
	})

	filteredUsers := make([]model.User, 0)
	for _, user := range users {
		if strings.EqualFold(user.LastName, "Faria") {
			filteredUsers = append(filteredUsers, user)
		}
	}

	return filteredUsers, nil
}

func (u *userService) ListUserLastNames(ctx context.Context) ([]string, error) {
	users, err := u.userRepository.ListUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("error listing users: %w", err)
	}

	userNames := make([]string, 0)
	for _, user := range users {
		userNames = append(userNames, user.LastName)
	}
	// first option
	userNames = removeDuplicates(userNames)
	// second option
	userNames = removeDuplicatesUsingMap(userNames)

	return userNames, nil
}

func removeDuplicates(elements []string) []string {
	encountered := map[string]bool{}
	var result []string

	for v := range elements {
		if encountered[elements[v]] {
			// Do not add duplicate.
		} else {
			encountered[elements[v]] = true
			result = append(result, elements[v])
		}
	}

	return result
}

func removeDuplicatesUsingMap(elements []string) []string {
	uniqueStrings := map[string]bool{}

	for v := range elements {
		uniqueStrings[elements[v]] = true
	}

	keys := make([]string, 0, len(uniqueStrings))
	for k := range uniqueStrings {
		keys = append(keys, k)
	}
	return keys
}
