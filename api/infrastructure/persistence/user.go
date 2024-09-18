package persistence

import (
	"context"

	"github.com/ei-sugimoto/tatekae/api/infrastructure"
	"github.com/ei-sugimoto/tatekae/api/model"
	"github.com/ei-sugimoto/tatekae/api/repo"
)

type PersistUser struct {
	db *infrastructure.DB
}

func NewPersistUser(db *infrastructure.DB) repo.UserRepo {
	return &PersistUser{
		db: db,
	}
}

func (p *PersistUser) Create(user *model.User) (*model.User, error) {
	ctx := context.Background()
	res, err := p.db.User.Create().SetUsername(user.Username).SetEmail(user.Email).SetPassword(user.Password).SetCreatedAt(user.CreateAt).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:       res.ID,
		Username: res.Username,
		Email:    res.Email,
		CreateAt: res.CreatedAt,
		Password: res.Password,
	}, nil
}
