package user

type UserUsecase struct {
	userRepository UserRepository
}

func NewUserUsecase(repo UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepository: repo,
	}
}
