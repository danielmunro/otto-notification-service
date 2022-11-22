package model

import (
	"encoding/json"
	"time"
)

type Post struct {
	Uuid       string     `json:"uuid"`
	Text       string     `json:"text,omitempty"`
	Draft      bool       `json:"draft"`
	Visibility Visibility `json:"visibility,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`
	User       User       `json:"user"`
}

func DecodeMessageToPost(message []byte) (*Post, error) {
	post := &Post{}
	err := json.Unmarshal(message, post)
	return post, err
}
