package persistence

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/ei-sugimoto/tatekae/api/infrastructure"
	"github.com/ei-sugimoto/tatekae/api/infrastructure/ent/user"
	"github.com/ei-sugimoto/tatekae/api/model"
	"github.com/ei-sugimoto/tatekae/api/pkg"
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

func (p *PersistUser) Get(id int) (*model.User, *pkg.CustomErr) {
	ctx := context.Background()
	res, err := p.db.User.Query().Where(user.ID(id)).Only(ctx)
	if err != nil {
		return nil, &pkg.CustomErr{
			Code: 500,
			Err:  err,
		}
	}

	if res == nil {
		return nil, &pkg.CustomErr{
			Code: 404,
			Err:  errors.New("user not found"),
		}
	}

	return &model.User{
		ID:       res.ID,
		Username: res.Username,
		Email:    res.Email,
		CreateAt: res.CreatedAt,
		Password: res.Password,
	}, nil
}

func (p *PersistUser) Login(email, password string) (*model.User, *pkg.CustomErr) {
	ctx := context.Background()
	res, err := p.db.User.Query().Where(user.And(user.Email(email), user.Password(password))).Only(ctx)
	if err != nil {
		return nil, &pkg.CustomErr{
			Code: 500,
			Err:  err,
		}
	}

	if errors.Is(err, sql.ErrNoRows) {
		return nil, &pkg.CustomErr{
			Code: 404,
			Err:  errors.New("user not found"),
		}
	}

	if res.Password != password {
		return nil, &pkg.CustomErr{
			Code: 401,
			Err:  errors.New("password is incorrect"),
		}
	}

	return &model.User{
		ID:       res.ID,
		Username: res.Username,
		Email:    res.Email,
		CreateAt: res.CreatedAt,
		Password: res.Password,
	}, nil
}

func (p *PersistUser) ListByIDs(ids []int) ([]*model.User, error) {
	ctx := context.Background()
	res, err := p.db.User.Query().Where(user.IDIn(ids...)).All(ctx)
	if err != nil {
		return nil, err
	}
	log.Println(res)

	users := make([]*model.User, len(res))
	for i, u := range res {
		users[i] = &model.User{
			ID:       u.ID,
			Username: u.Username,
			Email:    u.Email,
			CreateAt: u.CreatedAt,
			Password: u.Password,
		}
	}

	return users, nil
}
