package repository

import (
	"github.com/fahmijufri/jismi/domain"
)

type UserInterface interface {
	Save(domain.User) error
	FindOne(domain.User) (*domain.User, error)
	FindList(domain.User) ([]*domain.User, error)
}
