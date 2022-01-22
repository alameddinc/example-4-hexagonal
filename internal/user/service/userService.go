package service

import (
	"hexagonal/internal/user/domain"
	"hexagonal/internal/user/repository"
)

type Service interface {
	Create(int, string) error
	Update(int, string) error
	Delete(int) error
	Get(int) (*domain.User, error)
}

type userService struct {
	userRepo repository.Repository
}

// Service Creator
func NewUserService(userRepo repository.Repository) Service {
	return &userService{userRepo: userRepo}
}

func (s *userService) Create(id int, name string) error {
	return s.userRepo.Insert(id, name)
}
func (s *userService) Update(id int, name string) error {
	return s.userRepo.UpdateName(id, name)
}
func (s *userService) Delete(id int) error {
	return s.userRepo.Delete(id)
}
func (s *userService) Get(id int) (*domain.User, error) {
	tmpResponse, err := s.userRepo.Get(id)
	if err != nil {
		return nil, err
	}
	return &domain.User{
		ID:   tmpResponse["id"].(int),
		Name: tmpResponse["name"].(string),
	}, nil
}
