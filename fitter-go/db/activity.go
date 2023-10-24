package storage

import (
	"github.com/notaduck/fitter-go/models"
	"github.com/tormoder/fit"
)

func (s *PostgresStore) CreateActivity(userId int, activity *fit.ActivityFile) (int64, error) {

	var activityEntity *models.Activity

	var records []*models.Record

	for _, record := range activity.Records {

		records = append(records, models.Record{
			// Model:                         gorm.Model{},
			// ID:                            userId,
			// ActivityID:                    0,
			Timestamp:                     record.Timestamp,
			PositionLat:                   record.PositionLat.Degrees(),
			PositionLong:                  record.PositionLong.Degrees(),
			Altitude:                      0,
			HeartRate:                     0,
			Cadence:                       0,
			Distance:                      0,
			Speed:                         0,
			Power:                         0,
			CompressedSpeedDistance:       []byte{},
			Grade:                         0,
			Resistance:                    0,
			TimeFromCourse:                0,
			CycleLength:                   0,
			Temperature:                   0,
			Speed1s:                       []byte{},
			Cycles:                        0,
			TotalCycles:                   0,
			CompressedAccumulatedPower:    0,
			AccumulatedPower:              0,
			LeftRightBalance:              0,
			GPSAccuracy:                   0,
			VerticalSpeed:                 0,
			Calories:                      0,
			VerticalOscillation:           0,
			StanceTimePercent:             0,
			StanceTime:                    0,
			ActivityType:                  0,
			LeftTorqueEffectiveness:       0,
			RightTorqueEffectiveness:      0,
			LeftPedalSmoothness:           0,
			RightPedalSmoothness:          0,
			CombinedPedalSmoothness:       0,
			Time128:                       0,
			StrokeType:                    0,
			Zone:                          0,
			BallSpeed:                     0,
			Cadence256:                    0,
			FractionalCadence:             0,
			TotalHemoglobinConc:           0,
			TotalHemoglobinConcMin:        0,
			TotalHemoglobinConcMax:        0,
			SaturatedHemoglobinPercent:    0,
			SaturatedHemoglobinPercentMin: 0,
			SaturatedHemoglobinPercentMax: 0,
			DeviceIndex:                   0,
			EnhancedSpeed:                 0,
			EnhancedAltitude:              0,
		})

	}

	if result := s.db.Create(activityEntity); result.Error != nil {
		return -1, result.Error
	}

	return -1, nil
}

func (s *PostgresStore) CreateActivities(activities *[]models.Activity) error {

	result := s.db.Create(activities)

	if result.Error != nil {
		return result.Error
	}

	return nil

}

func (s *PostgresStore) GetActivities(userId int) ([]*models.Activity, error) {

	var activities []*models.Activity

	if err := s.db.Where("user_id = ?", userId).Scan(activities); err.Error != nil {
		return nil, err.Error
	}

	return activities, nil
}

func (s *PostgresStore) GetActivity(userId, activityId int) (*models.Activity, error) {

	var activity *models.Activity

	if err := s.db.Where("user_id = ? AND id = ?", userId).Scan(activity); err.Error != nil {
		return nil, err.Error
	}

	return activity, nil

}
