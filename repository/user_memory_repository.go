package repository

import (
	"web-api/entity"

	"github.com/google/uuid"
)

type UserMemoryRepository struct {
	users map[uuid.UUID]*entity.User
}

func NewUserMemoryRepository() *UserMemoryRepository {
	users := &UserMemoryRepository{
		users: make(map[uuid.UUID]*entity.User),
	}
	// Data mock - just for testing purpose
	u := &entity.User{
		ID:    uuid.MustParse("49bae575-85de-4fa6-86ac-c96d5d03d968"),
		Name:  "Felix",
		Email: "phong.danghong.dev@gmail.com",
	}
	users.users[u.ID] = u
	return users
}

func (r *UserMemoryRepository) GetUserByID(id uuid.UUID) (*entity.User, error) {
	user, ok := r.users[id]
	if !ok {
		return nil, nil
	}
	return user, nil
}

// Store stores a user.
func (r *UserMemoryRepository) Store(user *entity.User) error {
	r.users[user.ID] = user
	return nil
}
