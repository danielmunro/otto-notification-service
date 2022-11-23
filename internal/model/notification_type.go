package model

type NotificationType string

// List of NotificationType
const (
	POST_LIKED NotificationType = "post_liked"
	FOLLOWED   NotificationType = "followed"
	REPLIED    NotificationType = "replied"
)
