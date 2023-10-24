package storage

import (
	"time"

	"github.com/notaduck/fitter-go/models"
)

func (s *PostgresStore) CreateUser(user *models.User) error {

	userEntity := s.db.Create(&user)

	err := userEntity.Error

	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) GetUserByEmail(email string) (*models.User, error) {

	var user models.User

	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {

		return nil, err
	}

	return &user, nil
}

func (s *PostgresStore) GetUserById(id int) (*models.User, error) {

	var user models.User

	if err := s.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil

}

func NewUser(firstName, lastName, password string) (*models.User, error) {

	return &models.User{
		FirstName: firstName,
		LastName:  lastName,
		CreatedAt: time.Now().UTC(),
	}, nil

}

func (s *PostgresStore) GetUserByAuth0Id(auth0Id string) (*models.User, error) {

	var user models.User

	if err := s.db.Where("auth0_id = ?", auth0Id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil

}
