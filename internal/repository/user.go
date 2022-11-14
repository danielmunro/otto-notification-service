package repository

import (
	"errors"
	"github.com/danielmunro/otto-notification-service/internal/entity"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	conn *gorm.DB
}

func CreateUserRepository(conn *gorm.DB) *UserRepository {
	return &UserRepository{conn}
}

func (u *UserRepository) FindOneByUuid(uuid uuid.UUID) (*entity.User, error) {
	user := &entity.User{}
	u.conn.Where("uuid = ?", uuid).Find(user)
	if user.ID == 0 {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (u *UserRepository) Create(user *entity.User) {
	u.conn.Create(user)
}

func (u *UserRepository) Save(user *entity.User) {
	u.conn.Save(user)
}
