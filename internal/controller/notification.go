package controller

import (
	"encoding/json"
	"github.com/danielmunro/otto-notification-service/internal/model"
	"github.com/danielmunro/otto-notification-service/internal/service"
	"github.com/google/uuid"
	"net/http"
)

const notificationLimit = 100

// AcknowledgeNotificationsForUserV1 - Acknowledge notifications for a user
func AcknowledgeNotificationsForUserV1(w http.ResponseWriter, r *http.Request) {
	session := service.CreateDefaultAuthService().GetSessionFromRequest(r)
	if session == nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	notificationModel, _ := model.DecodeRequestToNotificationAcknowledgement(r)
	err := service.CreateNotificationService().AcknowledgeNotificationsForUser(
		uuid.MustParse(session.User.Uuid),
		notificationModel,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

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
