package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/notaduck/fitter-go/models"
	"github.com/tormoder/fit"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage interface {
	CreateActivities(activities *[]models.Activity) error
	CreateActivity(userId int, activity *fit.ActivityFile) (int64, error)
	GetActivities(userId int) ([]*Activity, error)
	GetActivity(userId, activityId int) (*Activity, error)

	CreateMsgRecords(records []*fit.RecordMsg, activityId int64) error
	GetRecordMsgs(activityId int) ([]*models.RecordDTO, error)
	CreateUser(*models.User) error
	GetUserById(id int) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByAuth0Id(auth0Id string) (*models.User, error)
}

type PostgresStore struct {
	db *gorm.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "host=db user=fitter dbname=fitter password=fitter sslmode=disable port=5432"

	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	// db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// if err := db.Ping(); err != nil {
	// 	return nil, err
	// }

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {

	if err := s.createUserTable(); err != nil {
		log.Fatal(err)
		return err
	}

	if err := s.createActivityTable(); err != nil {
		log.Fatal(err)
		return err
	}

	if err := s.createRecordMsgTable(); err != nil {
		log.Fatal(err)
		return err
	}

	return nil

}

func Transact(db *sql.DB, txFunc func(*sql.Tx) error) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	err = txFunc(tx)
	return err

}
