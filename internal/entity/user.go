package entity

import (
	"github.com/danielmunro/otto-notification-service/internal/model"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Uuid       *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Username   string     `gorm:"unique;not null"`
	ProfilePic string
	Name       string
	Role       string `gorm:"default:'user'"`
	IsBanned   bool   `gorm:"default:false"`
}

func (u *User) UpdateUserProfileFromModel(user *model.User) {
	u.Name = user.Name
	u.ProfilePic = user.ProfilePic
	u.Username = user.Username
	u.IsBanned = user.IsBanned
}
