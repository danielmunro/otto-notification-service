package entity

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Notification struct {
	gorm.Model
	Uuid              *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	UserID            uint       `gorm:"foreignkey:User"`
	User              *User
	Seen              bool
	Link              string
	Description       string
	TriggeredByUserID uint `gorm:"foreignkey:User"`
	TriggeredByUser   *User
}
