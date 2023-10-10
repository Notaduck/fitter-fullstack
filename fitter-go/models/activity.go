package models

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	ID             int       `gorm:"primaryKey" json:"id"`
	Timestamp      time.Time `json:"timeStamp"`
	TotalTimerTime int       `json:"totalTimerTime"`
	NumSessions    int       `json:"numSessions"`
	Type           int       `json:"type"`
	Event          int       `json:"event"`
	EventType      int       `json:"eventType"`
	LocalTimestamp time.Time `json:"localTimestamp"`
	EventGroup     int       `json:"eventGroup"`
}
