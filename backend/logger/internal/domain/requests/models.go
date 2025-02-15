package requests

import "time"

type LogPOST struct {
	UserID string    `json:"user_id"`
	Action string    `json:"type"`
	Time   time.Time `json:"time"`
	Info   string    `json:"info"`
}
