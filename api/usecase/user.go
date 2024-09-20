package usecase

import (
	"github.com/ei-sugimoto/tatekae/api/model"
	"github.com/ei-sugimoto/tatekae/api/pkg"
	"github.com/ei-sugimoto/tatekae/api/repo"
)

type IUserUsecase interface {
	Create(*model.User) (*model.User, error)
	Get(id int) (*model.User, *pkg.CustomErr)
	Login(string, string) (*model.User, *pkg.CustomErr)
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

func (u *UserUsecase) Get(id int) (*model.User, *pkg.CustomErr) {
	return u.repo.Get(id)
}

func (u *UserUsecase) Login(email, password string) (*model.User, *pkg.CustomErr) {
	return u.repo.Login(email, password)
}
