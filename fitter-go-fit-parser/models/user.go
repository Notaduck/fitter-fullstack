package models

import "time"

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
}
