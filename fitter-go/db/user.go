package storage

import (
	"time"

	"github.com/notaduck/fitter-go/models"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func (s *PostgresStore) createUserTable() error {
	query := `create table if not exists users (
		id serial primary key,
		first_name varchar(100),
		last_name varchar(100),
		encrypted_password varchar(100),
		email 	varchar(100),
		user_name 	varchar(100),
		created_at timestamp
	);`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateUser(user *models.User) error {
	query := `insert into users 
	(first_name, last_name, encrypted_password, created_at, email, user_name)
	values ($1, $2, $3, $4, $5, $6)`

	_, err := s.db.Query(
		query,
		user.FirstName,
		user.LastName,
		user.EncryptedPassword,
		user.CreatedAt,
		user.Email,
		user.Username,
	)

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
