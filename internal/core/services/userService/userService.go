package userService

import (
	"hexagonal/internal/core/domains"
	ports2 "hexagonal/internal/core/ports"
)

type userService struct {
	userRepo ports2.UserRepository
}

//NewUserService Creator
func NewUserService(userRepo ports2.UserRepository) ports2.UserService {
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
func (s *userService) Get(id int) (*domains.User, error) {
	tmpResponse, err := s.userRepo.Get(id)
	if err != nil {
		return nil, err
	}
	return &domains.User{
		ID:   tmpResponse["id"].(int),
		Name: tmpResponse["name"].(string),
	}, nil
}
