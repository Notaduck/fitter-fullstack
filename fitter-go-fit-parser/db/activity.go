package storage

import (
	"database/sql"
	"fmt"
	"math"

	"github.com/notaduck/fitter-go-fit-parser/models"
	"github.com/tormoder/fit"
)

func (s *PostgresStore) CreateActivity(userId int, activity *fit.ActivityFile) (int64, error) {

	var id int64

	err := Transact(s.db, func(tx *sql.Tx) error {

		var filteredRedords = []*fit.RecordMsg{}
		var distance float64

		for index, record := range activity.Records {
			if !(record.PositionLat.Invalid() && record.PositionLong.Invalid()) {
				filteredRedords = append(filteredRedords, record)

				if index != 0 {

					lat2 := record.PositionLat.Degrees()

					long2 := record.PositionLong.Degrees()

					lat1 := activity.Records[index-1].PositionLat.Degrees()

					long1 := activity.Records[index-1].PositionLong.Degrees()
					if !math.IsNaN(lat1) && !math.IsNaN(long1) && !math.IsNaN(lat2) && !math.IsNaN(long2) {

						distance = distance + haversine(lat1, long1, lat2, long2)
					}

				}

			}

		}
		fmt.Println("D", distance)
		fmt.Printf("Distance: %.2f km\n", distance)
		// fmt.Println("avg_speed", avgSpeed)
		// fmt.Println("total_distance", totalDistance)

		stmt, err := tx.Prepare(`
            INSERT INTO activities (
				user_id,
                timestamp,
                total_timer_time,
                num_sessions,
                type,
                event,
                event_type,
                local_timestamp,
                event_group,
				distance
            ) VALUES (
                $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
            ) RETURNING id
        `)
		if err != nil {
			return err
		}
		defer stmt.Close()

		err = stmt.QueryRow(
			userId,
			activity.Activity.Timestamp,
			activity.Activity.TotalTimerTime,
			activity.Activity.NumSessions,
			activity.Activity.Type,
			activity.Activity.Event,
			activity.Activity.EventType,
			activity.Activity.LocalTimestamp,
			activity.Activity.EventGroup,
			distance,
		).Scan(&id)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return -1, err
	}

	return id, nil
}

func (s *PostgresStore) CreateActivities(activities *[]models.Activity) error {

	return Transact(s.db, func(tx *sql.Tx) error {

		stmt, err := tx.Prepare(`
            INSERT INTO activities (
                timestamp,
                total_timer_time,
                num_sessions,
                type,
                event,
                event_type,
                local_timestamp,
                event_group
            ) VALUES (
                $1, $2, $3, $4, $5, $6, $7, $8
            ) RETURNING id
        `)

		if err != nil {
			panic(err)
		}

		defer stmt.Close()
		var id int64
		for _, activity := range *activities {
			err := stmt.QueryRow(
				activity.Timestamp,
				activity.TotalTimerTime,
				activity.NumSessions,
				activity.Type,
				activity.Event,
				activity.EventType,
				activity.LocalTimestamp,
				activity.EventGroup,
			).Scan(&id)

			if err != nil {
				// An error occurred, panic to trigger rollback
				panic(err)
			}
		}
		fmt.Println(id)

		return nil
	})
}

func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	// Convert degrees to radians
	lat1 = degToRad(lat1)
	lon1 = degToRad(lon1)
	lat2 = degToRad(lat2)
	lon2 = degToRad(lon2)

	// Haversine formula
	dlat := lat2 - lat1
	dlon := lon2 - lon1
	a := math.Pow(math.Sin(dlat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(dlon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	r := 6371.0 // Earth's radius in km
	println(a, c, r)
	return r * c
}

func degToRad(deg float64) float64 {
	return deg * (math.Pi / 180)
}
