package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type CreateUserDTO struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Username  string `json:"username"`
}

type User struct {
	ID                int       `json:"id"`
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	EncryptedPassword string    `json:"-"`
	Email             string    `json:"email"`
	Username          string    `json:"username"`
	CreatedAt         time.Time `json:"createdAt"`
	IssuerURL         time.Time `json:"issuerUrl"`
	Auth0ID           string    `json:"auth0_id"`
}

func NewUser(firstName, lastName, password, username, email string) (*User, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		FirstName:         firstName,
		LastName:          lastName,
		EncryptedPassword: string(encpw),
		Email:             email,
		Username:          username,
		CreatedAt:         time.Now().UTC(),
	}, nil
}

func (u *User) ValidPassword(pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(pw)) == nil
}
