package model

import (
	"encoding/json"
	"net/http"
	"time"
)

type NotificationAcknowledgement struct {
	DatetimeAcknowledged time.Time `json:"datetime_acknowledged"`
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
