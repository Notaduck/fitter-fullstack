package models

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	ID             int       `gorm:"primaryKey" json:"id"`
	Distance       float64   `json:"distance"`
	Elevation      float32   `json:"elevation"`
	TotalRideTime  int       `json:"totalRideTime"`
	Timestamp      time.Time `json:"timeStamp"`
	TotalTimerTime int       `json:"totalTimerTime"`
	NumSessions    int       `json:"numSessions"`
	Type           int       `json:"type"`
	Event          int       `json:"event"`
	EventType      int       `json:"eventType"`
	LocalTimestamp time.Time `json:"localTimestamp"`
	EventGroup     int       `json:"eventGroup"`
	UserId         int       `json:"userId"`
	Records        []Record  `json:"records"`
}
