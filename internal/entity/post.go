package entity

import (
	"github.com/danielmunro/otto-notification-service/internal/model"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Text       string
	Draft      bool
	UserID     uint
	User       *User
	Visibility model.Visibility `gorm:"default:'public'"`
	Uuid       *uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4()"`
}

func (p *Post) UpdatePostFromModel(post *model.Post) {
	p.Text = post.Text
	p.Visibility = post.Visibility
	p.Draft = post.Draft
}
