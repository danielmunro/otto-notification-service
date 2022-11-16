package model

import (
	"encoding/json"
	"net/http"
	"time"
)

type NotificationAcknowledgement struct {
	DatetimeStarted time.Time `json:"datetime_started"`
	DatetimeEnded   time.Time `json:"datetime_ended"`
}

func DecodeRequestToNotificationAcknowledgement(r *http.Request) (*NotificationAcknowledgement, error) {
	decoder := json.NewDecoder(r.Body)
	var data *NotificationAcknowledgement
	err := decoder.Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
