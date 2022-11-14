package model

type Notification struct {
	Uuid            string `json:"uuid"`
	User            User   `json:"user"`
	Seen            bool   `json:"seen"`
	Link            string `json:"link"`
	Description     string `json:"description"`
	TriggeredByUser User   `json:"triggered_by_user,omitempty"`
}
