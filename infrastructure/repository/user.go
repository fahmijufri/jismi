package repository

import (
	"github.com/fahmijufri/jismi/domain"

	"github.com/jinzhu/gorm"
)

const UserTableName = "users"

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u UserRepository) Save(user domain.User) error {
	return u.db.
		Omit("created_at").
		Table(UserTableName).
		Save(&user).Error
}

func (u UserRepository) FindOne(query domain.User) (*domain.User, error) {
	var (
		result domain.User
	)
	err := u.db.
		Table(UserTableName).
		Where(query).
		First(&result).
		Error
	return &result, err
}

func (u UserRepository) FindList(query domain.User) ([]*domain.User, error) {
	var (
		result []*domain.User
	)
	err := u.db.
		Table(UserTableName).
		Where(query).
		Find(&result).
		Error
	return result, err
}
