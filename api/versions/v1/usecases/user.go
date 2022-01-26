package usecases

import (
	"github.com/aofdev/fiber-versioning-boilerplate/api/versions/v1/entities"
	repositories "github.com/aofdev/fiber-versioning-boilerplate/api/versions/v1/repositories"
)

type IUserUsecase interface {
	GetAllUser() (*entities.User, error)
}

type UserUsecase struct {
	userRepo repositories.IUserRepository
}

func NewUserUsecase(repo repositories.IUserRepository) IUserUsecase {
	return &UserUsecase{
		userRepo: repo,
	}
}

func (u *UserUsecase) GetAllUser() (*entities.User, error) {
	// business logic here..
	id := "1"
	return u.userRepo.GetAllUser(id)
}
