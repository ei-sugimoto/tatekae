package usecase

import (
	"github.com/ei-sugimoto/tatekae/api/model"
	"github.com/ei-sugimoto/tatekae/api/repo"
)

type IUserUsecase interface {
	Create(*model.User) (*model.User, error)
}

type UserUsecase struct {
	repo repo.UserRepo
}

func NewUserUsecase(repo repo.UserRepo) IUserUsecase {
	return &UserUsecase{
		repo: repo,
	}
}

func (u *UserUsecase) Create(user *model.User) (*model.User, error) {
	return u.repo.Create(user)
}
