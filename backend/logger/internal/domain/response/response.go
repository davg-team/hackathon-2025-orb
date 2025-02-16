package response

import (
	"time"
)

type LogsResponse struct {
	ID     string       `json:"id"`
	Action string       `json:"type"`
	Time   time.Time    `json:"time"`
	Info   string       `json:"info"`
	User   UserMetadata `json:"user"`
}

type UserMetadata struct {
	ID           int    `json:"id"`
	FullName     string `json:"full_name"`
	Organization string `json:"organization"`
	Email        string `json:"email"`
	Role         string `json:"role"`
}
