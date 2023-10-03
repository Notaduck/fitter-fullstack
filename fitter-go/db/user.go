package storage

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/notaduck/fitter-go/models"
	"golang.org/x/crypto/bcrypt"
)

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

	rows, err := s.db.Query("select * from users where email = $1", email)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoUser(rows)
	}

	return nil, fmt.Errorf("user %s not found", email)
}

func (s *PostgresStore) GetUserById(id int) (*models.User, error) {
	rows, err := s.db.Query("select * from users where id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoUser(rows)
	}

	return nil, fmt.Errorf("user %d not found", id)
}

func NewUser(firstName, lastName, password string) (*models.User, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &models.User{
		FirstName:         firstName,
		LastName:          lastName,
		EncryptedPassword: string(encpw),
		CreatedAt:         time.Now().UTC(),
	}, nil
}
func scanIntoUser(rows *sql.Rows) (*models.User, error) {
	user := new(models.User)
	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.EncryptedPassword,
		&user.Email,
		&user.Username,
		&user.CreatedAt)

	return user, err
}
