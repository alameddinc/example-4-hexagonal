package ports

import (
	"hexagonal/internal/core/domains"
)

type UserService interface {
	Create(int, string) error
	Update(int, string) error
	Delete(int) error
	Get(int) (*domains.User, error)
}

