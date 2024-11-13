// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package repository

import (
	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `json:"id"`
	Email          string    `json:"email"`
	Username       string    `json:"username"`
	HashedPassword []byte    `json:"hashed_password"`
	Salt           []byte    `json:"salt"`
	Active         bool      `json:"active"`
	Admin          bool      `json:"admin"`
}