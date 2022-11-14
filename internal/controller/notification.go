package controller

import (
	"encoding/json"
	"github.com/danielmunro/otto-notification-service/internal/service"
	"github.com/google/uuid"
	"net/http"
)

const notificationLimit = 100

// GetNotificationsForUserV1 - Get notifications for a user
func GetNotificationsForUserV1(w http.ResponseWriter, r *http.Request) {
	session := service.CreateDefaultAuthService().GetSessionFromRequest(r)
	if session == nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	notifications, err := service.CreateNotificationService().
		GetNotificationsForUser(uuid.MustParse(session.User.Uuid), notificationLimit)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data, err := json.Marshal(notifications)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(data)
}
