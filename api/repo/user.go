package repo

import (
	"github.com/ei-sugimoto/tatekae/api/model"
	"github.com/ei-sugimoto/tatekae/api/pkg"
)

type UserRepo interface {
	Create(*model.User) (*model.User, error)
	Get(int) (*model.User, *pkg.CustomErr)
	Login(string, string) (*model.User, *pkg.CustomErr)
}
