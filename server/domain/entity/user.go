package entity

import "github.com/google/uuid"

type User struct {
	UUID     uuid.UUID
	Name     string
	Email    string
	Password string
}
