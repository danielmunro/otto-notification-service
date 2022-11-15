package model

type NotificationType string

// List of NotificationType
const (
	LIKED    NotificationType = "liked"
	FOLLOWED NotificationType = "followed"
	REPLIED  NotificationType = "replied"
)
