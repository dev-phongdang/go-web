package entity

import "github.com/google/uuid"

// User represents a user.
type User struct {
	ID       uuid.UUID
	Name     string
	Email    string
	Password string
}
