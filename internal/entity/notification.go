package entity

import (
	"github.com/danielmunro/otto-notification-service/internal/model"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Notification struct {
	gorm.Model
	Uuid              *uuid.UUID             `gorm:"type:uuid;default:uuid_generate_v4()"`
	Seen              bool                   `gorm:"default:false"`
	Link              string                 `gorm:"not null"`
	NotificationType  model.NotificationType `gorm:"unique_index:idx_user_triggered_type"`
	UserID            uint                   `gorm:"unique_index:idx_user_triggered_type"`
	User              *User
	TriggeredByUserID uint `gorm:"unique_index:idx_user_triggered_type"`
	TriggeredByUser   *User
}
