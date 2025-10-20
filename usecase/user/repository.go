package user

import (
	"web-api/entity"

	"github.com/google/uuid"
)

type UserRepository interface {
	GetUserByID(id uuid.UUID) (*entity.User, error)
	Store(user *entity.User) error
}
