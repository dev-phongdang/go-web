package user

import (
	"web-api/entity"
)

// Store stores a user.
func (uc *UserUsecase) Store(user *entity.User) error {

	err := uc.userRepository.Store(user)
	if err != nil {
		return err
	}
	return nil
}
