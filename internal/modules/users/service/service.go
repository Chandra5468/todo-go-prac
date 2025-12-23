package service

import (
	"errors"

	"github.com/Chandra5468/todo-go/internal/modules/users"
)

// If you want u can separate repository and service logic if needed

type Service interface {
	GetUserDetails(id int) (users.User, error)
}

type svc struct {
	repo     users.UserRepository
	notifier users.NotificationSender
}

func NewService(repo users.UserRepository, notifier users.NotificationSender) Service {
	return &svc{
		repo:     repo,
		notifier: notifier,
	}
}

func (s *svc) GetUserDetails(id int) (users.User, error) {

	// Business Logic: Check if user exists
	user, err := s.repo.FindByID(id)
	if err != nil {
		return users.User{}, err
	}
	if user.Username == "" {
		return users.User{}, errors.New("user not found")
	}

	return user, nil

}
