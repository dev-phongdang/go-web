package user

type UserUsecase struct {
	userRepository UserRepository
}

// NewUserUsecase creates a new user use case.
func NewUserUsecase(repo UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepository: repo,
	}
}
