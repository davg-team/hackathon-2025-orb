package domain

import "time"

type LogModel struct {
	ID     string    `bson:"_id" json:"id"`
	UserID string    `bson:"user_id" json:"user_id"`
	Action string    `bson:"type" json:"type"`
	Time   time.Time `bson:"time" json:"time"`
	Info   string    `bson:"info" json:"info"`
}
