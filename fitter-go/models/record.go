package models

import "gorm.io/gorm"

type Record struct {
	gorm.Model
	ID                            int `gorm:"primaryKey" json:"id"`
	ActivityID                    int
	Timestamp                     string
	PositionLat                   float64
	PositionLong                  float64
	Altitude                      int16
	HeartRate                     int16
	Cadence                       int16
	Distance                      int
	Speed                         int16
	Power                         int16
	CompressedSpeedDistance       []byte
	Grade                         int16
	Resistance                    int16
	TimeFromCourse                int
	CycleLength                   int16
	Temperature                   int16
	Speed1s                       []byte
	Cycles                        int16
	TotalCycles                   int
	CompressedAccumulatedPower    int16
	AccumulatedPower              int
	LeftRightBalance              int16
	GPSAccuracy                   int16
	VerticalSpeed                 int16
	Calories                      int16
	VerticalOscillation           int16
	StanceTimePercent             int16
	StanceTime                    int16
	ActivityType                  int16
	LeftTorqueEffectiveness       int16
	RightTorqueEffectiveness      int16
	LeftPedalSmoothness           int16
	RightPedalSmoothness          int16
	CombinedPedalSmoothness       int16
	Time128                       int16
	StrokeType                    int16
	Zone                          int16
	BallSpeed                     int16
	Cadence256                    int16
	FractionalCadence             int16
	TotalHemoglobinConc           int16
	TotalHemoglobinConcMin        int16
	TotalHemoglobinConcMax        int16
	SaturatedHemoglobinPercent    int16
	SaturatedHemoglobinPercentMin int16
	SaturatedHemoglobinPercentMax int16
	DeviceIndex                   int16
	EnhancedSpeed                 int
	EnhancedAltitude              int
}