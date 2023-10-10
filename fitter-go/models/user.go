package models

import (
	"time"

	"gorm.io/gorm"
)

type CreateUserDTO struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Username  string `json:"username"`
}

type User struct {
	gorm.Model
	ID        int       `gorm:"primaryKey" json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `gorm:"unique" json:"email"`
	Username  string    `gorm:"unique" json:"username"`
	CreatedAt time.Time `json:"createdAt"`
	IssuerURL string    `json:"issuerUrl"`
	Auth0ID   string    `gorm:"unique" json:"auth0_id"`
}

func NewUser(firstName, lastName, password, username, email string) (*User, error) {

	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Username:  username,
		CreatedAt: time.Now().UTC(),
	}, nil
}
