package models

import "time"

type RecordDTO struct {
	ID    int     `json:"id"`
	Lon   float64 `json:"lon"`
	Lat   float64 `json:"lat"`
	Speed int     `json:"speed"`
}

type GetActivityDTO struct {
	ID            int          `json:"id"`
	Timestamp     time.Time    `json:"timestamp"`
	TotalRideTime int          `json:"totalRideTime"`
	Distance      float64      `json:"distance"`
	Elevation     float32      `json:"elevation"`
	Records       []*RecordDTO `json:"records"`
}
