package user

import (
	"web-api/entity"

	"github.com/google/uuid"
)

// UserRepository is an interface for user repository.
type UserRepository interface {
	GetUserByID(id uuid.UUID) (*entity.User, error)
	Store(user *entity.User) error
}
