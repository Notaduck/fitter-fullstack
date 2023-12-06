package storage

// import "github.com/tormoder/fit"

// func (s *PostgresStore) CreateMsgRecords(records []*fit.RecordMsg, activityId int64) error {

// 	s.db.Create()

// }

// func (s *PostgresStore) CreateMsgRecords(records []*fit.RecordMsg, activityId int64) error {

// 	return Transact(s.db, func(tx *sql.Tx) error {
// 		stmt, err := s.db.Prepare(`
//             INSERT INTO record_msgs (
//                 activity_id,
//                 timestamp,
//                 position_lat,
//                 position_long,
//                 altitude,
//                 heart_rate,
//                 cadence,
//                 distance,
//                 speed,
//                 power,
//                 compressed_speed_distance,
//                 grade,
//                 resistance,
//                 time_from_course,
//                 cycle_length,
//                 temperature,
//                 speed_1s,
//                 cycles,
//                 total_cycles,
//                 compressed_accumulated_power,
//                 accumulated_power,
//                 left_right_balance,
//                 gps_accuracy,
//                 vertical_speed,
//                 calories,
//                 vertical_oscillation,
//                 stance_time_percent,
//                 stance_time,
//                 activity_type,
//                 left_torque_effectiveness,
//                 right_torque_effectiveness,
//                 left_pedal_smoothness,
//                 right_pedal_smoothness,
//                 combined_pedal_smoothness,
//                 time_128,
//                 stroke_type,
//                 zone,
//                 ball_speed,
//                 cadence_256,
//                 fractional_cadence,
//                 total_hemoglobin_conc,
//                 total_hemoglobin_conc_min,
//                 total_hemoglobin_conc_max,
//                 saturated_hemoglobin_percent,
//                 saturated_hemoglobin_percent_min,
//                 saturated_hemoglobin_percent_max,
//                 device_index,
//                 enhanced_speed,
//                 enhanced_altitude
//             ) VALUES (
//                 $1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
//                 $11, $12, $13, $14, $15, $16, $17, $18, $19, $20,
//                 $21, $22, $23, $24, $25, $26, $27, $28, $29, $30,
//                 $31, $32, $33, $34, $35, $36, $37, $38, $39, $40,
//                 $41, $42, $43, $44, $45, $46, $47, $48, $49

//             )`)

// 		if err != nil {
// 			panic(err)
// 		}

// 		for _, record := range records {

// 			_, err := stmt.Exec(
// 				activityId,
// 				record.Timestamp,
// 				record.PositionLat.String(),
// 				record.PositionLong.String(),
// 				record.Altitude,
// 				record.HeartRate,
// 				record.Cadence,
// 				record.Distance,
// 				record.Speed,
// 				record.Power,
// 				record.CompressedSpeedDistance,
// 				record.Grade,
// 				record.Distance,
// 				record.TimeFromCourse,
// 				record.CycleLength,
// 				record.Temperature,
// 				record.Speed1s,
// 				record.Cycles,
// 				record.TotalCycles,
// 				record.CompressedAccumulatedPower,
// 				record.AccumulatedPower,
// 				record.LeftRightBalance,
// 				record.GpsAccuracy,
// 				record.VerticalSpeed,
// 				record.Calories,
// 				record.VerticalOscillation,
// 				record.StanceTimePercent,
// 				record.StanceTime,
// 				record.ActivityType,
// 				record.LeftTorqueEffectiveness,
// 				record.RightTorqueEffectiveness,
// 				record.LeftPedalSmoothness,
// 				record.RightPedalSmoothness,
// 				record.CombinedPedalSmoothness,
// 				record.Time128,
// 				record.StrokeType,
// 				record.Zone,
// 				record.BallSpeed,
// 				record.Cadence256,
// 				record.FractionalCadence,
// 				record.TotalHemoglobinConc,
// 				record.TotalHemoglobinConcMin,
// 				record.TotalHemoglobinConcMax,
// 				record.SaturatedHemoglobinPercent,
// 				record.SaturatedHemoglobinPercentMin,
// 				record.SaturatedHemoglobinPercentMax,
// 				record.DeviceIndex,
// 				record.EnhancedSpeed,
// 				record.EnhancedAltitude,
// 			)
// 			if err != nil {
// 				// An error occurred, panic to trigger rollback
// 				panic(err)
// 			}
// 		}

// 		return nil
// 	})
// }
// func (s *PostgresStore) GetRecordMsgs(activityId int) ([]*models.RecordDTO, error) {

// 	var records *models

// 	rows, err := s.db.Query(`SELECT
// 								id,
// 								position_lat,
// 								position_long,
// 								speed
// 							 FROM
// 							 	record_msgs
// 						     WHERE
// 							 	activity_id = $1
// 	`, activityId)

// 	records := []*models.RecordDTO{}

// 	for rows.Next() {
// 		record, err := scanIntoRecordMsg(rows)
// 		if err != nil {
// 			return nil, err
// 		}
// 		records = append(records, record)
// 	}

// 	if err != nil {
// 		fmt.Print(err)
// 	}

// 	return records, nil
// }

// func scanIntoRecordMsg(rows *sql.Rows) (*models.RecordDTO, error) {
// 	record := new(models.RecordDTO)
// 	err := rows.Scan(
// 		&record.ID,
// 		&record.Lat,
// 		&record.Lon,
// 		&record.Speed,
// 	)

// 	return record, err
// }

// type RecordMsg struct {
// 	ID    int
// 	Long  float64
// 	Lat   float64
// 	Speed float64
// }
