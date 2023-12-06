package storage

import (
	_ "github.com/lib/pq"
	"github.com/notaduck/fitter-go/models"
	"github.com/tormoder/fit"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Storage interface {
	CreateActivities(activities *[]models.Activity) error
	CreateActivity(userId int, activity *fit.ActivityFile) (int64, error)
	GetActivities(userId int) ([]*models.Activity, error)
	GetActivity(userId, activityId int) (*models.Activity, error)

	// CreateMsgRecords(records []*fit.RecordMsg, activityId int64) error
	// GetRecordMsgs(activityId int) ([]*models.RecordDTO, error)

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

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		CreateBatchSize: 1000,
		Logger:          logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {

	if err := s.db.AutoMigrate(&models.User{}); err != nil {
		return err
	}

	if err := s.db.AutoMigrate(&models.Activity{}); err != nil {
		return err
	}

	if err := s.db.AutoMigrate(&models.Record{}); err != nil {
		return err
	}

	return nil

}
