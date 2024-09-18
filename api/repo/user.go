package repo

import "github.com/ei-sugimoto/tatekae/api/model"

type UserRepo interface {
	Create(*model.User) (*model.User, error)
}
