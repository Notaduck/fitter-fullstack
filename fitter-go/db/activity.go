package storage

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/notaduck/fitter-go/models"
	"github.com/tormoder/fit"
)

func (s *PostgresStore) createActivityTable() error {

	query := `CREATE TABLE IF NOT EXISTS activities (
    id SERIAL PRIMARY KEY,
	user_id				      INTEGER NOT NULL,
    timestamp TIMESTAMPTZ,
    total_timer_time INT,
    num_sessions INT,
    type INT,
    event INT,
    event_type INT,
    local_timestamp TIMESTAMPTZ,
    event_group INT,
	distance DOUBLE PRECISION,
	total_ride_time BIGINT,
	elevation bigint,
	CONSTRAINT fk_user
      FOREIGN KEY(user_id) 
	  REFERENCES users(id)
      ON DELETE CASCADE
);`

	_, err := s.db.Exec(query)

	return err
}

func (s *PostgresStore) CreateActivity(userId int, activity *fit.ActivityFile) (int64, error) {

	var id int64

	err := Transact(s.db, func(tx *sql.Tx) error {
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
                event_group
            ) VALUES (
                $1, $2, $3, $4, $5, $6, $7, $8, $9
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

func (s *PostgresStore) GetActivities(userId int) ([]*Activity, error) {

	rows, err := s.db.Query(`SELECT id,
								timestamp,
								total_timer_time,
								ROUND(distance::numeric,2),
								elevation

							FROM 
								activities
							WHERE user_id = $1
	`, userId)

	activities := []*Activity{}

	for rows.Next() {
		activity, err := scanIntoActivity(rows)
		if err != nil {
			return nil, err
		}
		activities = append(activities, activity)
	}

	if err != nil {
		fmt.Print(err)
	}

	return activities, nil
}

func (s *PostgresStore) GetActivity(userId, activityId int) (*Activity, error) {

	rows, err := s.db.Query(`SELECT a.id,
								a.timestamp,
								a.total_timer_time,
								ROUND(a.distance::numeric,2),
								a.elevation
							FROM activities a
							WHERE a.user_id = $1 AND a.id = $2
	`, userId, activityId)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close() // Close the rows when you're done with them

	if !rows.Next() {
		// No rows were returned by the query
		return nil, fmt.Errorf("No activity was found with for the user with the give activity id")
	}

	activity, err := scanIntoActivity(rows)

	if err != nil {
		fmt.Print(err)
	}

	return activity, nil
}

func (s *PostgresStore) GetAc(userId int) ([]*Activity, error) {

	rows, err := s.db.Query(`SELECT id,
								timestamp,
								total_timer_time,
								ROUND(distance::numeric,2),
								elevation

							FROM 
								activities
							WHERE user_id = $1
	`, userId)

	activities := []*Activity{}

	for rows.Next() {
		activity, err := scanIntoActivity(rows)
		if err != nil {
			return nil, err
		}
		activities = append(activities, activity)
	}

	if err != nil {
		fmt.Print(err)
	}

	return activities, nil
}

func scanIntoActivity(rows *sql.Rows) (*Activity, error) {

	var elevation sql.NullFloat64

	activity := new(Activity)
	err := rows.Scan(
		&activity.ID,
		&activity.Timestamp,
		&activity.TotalRideTime,
		&activity.Distance,
		&elevation,
	)

	if elevation.Valid {
		activity.Elevation = float32(elevation.Float64)
	} else {
		activity.Elevation = 0.0
	}

	return activity, err
}

type Activity struct {
	ID            int       `json:"id"`
	Timestamp     time.Time `json:"timestamp"`
	TotalRideTime int       `json:"totalRideTime"`
	Distance      float64   `json:"distance"`
	Elevation     float32   `json:"elevation"`
}

type GetActivitiesDTO struct {
	ID        int       `json:"id"`
	Timestamp time.Time `json:"timeStamp"`
}
