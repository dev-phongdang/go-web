package user

import (
	"web-api/entity"

	"github.com/google/uuid"
)

// GetUserByID retrieves a user by their ID.
func (uc *UserUsecase) GetUserByID(id uuid.UUID) (*entity.User, error) {
	u, err := uc.userRepository.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return u, nil
}
