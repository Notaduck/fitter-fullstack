package models

import (
	"time"

	"gorm.io/gorm"
)

type Record struct {
	gorm.Model
	ID                            int       `gorm:"column:id;primary_key"`
	ActivityID                    int       `gorm:"column:activity_id"`
	AccumulatedPower              int64     `gorm:"column:accumulated_power"`
	ActivityType                  int       `gorm:"column:activity_type"`
	Altitude                      int       `gorm:"column:altitude"`
	BallSpeed                     int64     `gorm:"column:ball_speed"`
	Cadence                       int       `gorm:"column:cadence"`
	Cadence256                    int       `gorm:"column:cadence_256"`
	Calories                      int       `gorm:"column:calories"`
	CombinedPedalSmoothness       int       `gorm:"column:combined_pedal_smoothness"`
	CompressedAccumulatedPower    int64     `gorm:"column:compressed_accumulated_power"`
	CompressedSpeedDistance       []byte    `gorm:"column:compressed_speed_distance"`
	CycleLength                   int       `gorm:"column:cycle_length"`
	Cycles                        int       `gorm:"column:cycles"`
	DeviceIndex                   int       `gorm:"column:device_index"`
	Distance                      int       `gorm:"column:distance"`
	EnhancedAltitude              int64     `gorm:"column:enhanced_altitude"`
	EnhancedSpeed                 int64     `gorm:"column:enhanced_speed"`
	FractionalCadence             int       `gorm:"column:fractional_cadence"`
	GPSAccuracy                   int       `gorm:"column:gps_accuracy"`
	Grade                         int       `gorm:"column:grade"`
	HeartRate                     int       `gorm:"column:heart_rate"`
	LeftPedalSmoothness           int       `gorm:"column:left_pedal_smoothness"`
	LeftRightBalance              int       `gorm:"column:left_right_balance"`
	LeftTorqueEffectiveness       int       `gorm:"column:left_torque_effectiveness"`
	PositionLat                   float64   `gorm:"column:position_lat"`
	PositionLong                  float64   `gorm:"column:position_long"`
	Power                         int       `gorm:"column:power"`
	Resistance                    int       `gorm:"column:resistance"`
	RightPedalSmoothness          int       `gorm:"column:right_pedal_smoothness"`
	RightTorqueEffectiveness      int       `gorm:"column:right_torque_effectiveness"`
	SaturatedHemoglobinPercent    int       `gorm:"column:saturated_hemoglobin_percent"`
	SaturatedHemoglobinPercentMax int       `gorm:"column:saturated_hemoglobin_percent_max"`
	SaturatedHemoglobinPercentMin int       `gorm:"column:saturated_hemoglobin_percent_min"`
	Speed                         int       `gorm:"column:speed"`
	Speed1s                       []byte    `gorm:"column:speed_1s"`
	StanceTime                    int       `gorm:"column:stance_time"`
	StanceTimePercent             int       `gorm:"column:stance_time_percent"`
	StrokeType                    int       `gorm:"column:stroke_type"`
	Temperature                   int       `gorm:"column:temperature"`
	Time128                       int       `gorm:"column:time_128"`
	TimeFromCourse                int       `gorm:"column:time_from_course"`
	Timestamp                     time.Time `gorm:"column:timestamp"`
	TotalCycles                   int64     `gorm:"column:total_cycles"`
	TotalHemoglobinConc           int       `gorm:"column:total_hemoglobin_conc"`
	TotalHemoglobinConcMax        int       `gorm:"column:total_hemoglobin_conc_max"`
	TotalHemoglobinConcMin        int       `gorm:"column:total_hemoglobin_conc_min"`
	VerticalOscillation           int       `gorm:"column:vertical_oscillation"`
	VerticalSpeed                 int       `gorm:"column:vertical_speed"`
	Zone                          int       `gorm:"column:zone"`

	ActivityId int `gorm:column:activityId`
	// Add more struct fields for constraints if needed
}

// type Record struct {
// 	gorm.Model
// 	ID                            int `gorm:"primaryKey" json:"id"`
// 	ActivityID                    int
// 	Timestamp                     string
// 	PositionLat                   float64
// 	PositionLong                  float64
// 	Altitude                      int16
// 	HeartRate                     int16
// 	Cadence                       int16
// 	Distance                      int
// 	Speed                         int16
// 	Power                         int16
// 	CompressedSpeedDistance       []byte
// 	Grade                         int16
// 	Resistance                    int16
// 	TimeFromCourse                int
// 	CycleLength                   int16
// 	Temperature                   int16
// 	Speed1s                       []byte
// 	Cycles                        int16
// 	TotalCycles                   int
// 	CompressedAccumulatedPower    int16
// 	AccumulatedPower              int
// 	LeftRightBalance              int16
// 	GPSAccuracy                   int16
// 	VerticalSpeed                 int16
// 	Calories                      int16
// 	VerticalOscillation           int16
// 	StanceTimePercent             int16
// 	StanceTime                    int16
// 	ActivityType                  int16
// 	LeftTorqueEffectiveness       int16
// 	RightTorqueEffectiveness      int16
// 	LeftPedalSmoothness           int16
// 	RightPedalSmoothness          int16
// 	CombinedPedalSmoothness       int16
// 	Time128                       int16
// 	StrokeType                    int16
// 	Zone                          int16
// 	BallSpeed                     int16
// 	Cadence256                    int16
// 	FractionalCadence             int16
// 	TotalHemoglobinConc           int16
// 	TotalHemoglobinConcMin        int16
// 	TotalHemoglobinConcMax        int16
// 	SaturatedHemoglobinPercent    int16
// 	SaturatedHemoglobinPercentMin int16
// 	SaturatedHemoglobinPercentMax int16
// 	DeviceIndex                   int16
// 	EnhancedSpeed                 int
// 	EnhancedAltitude              int
// }
