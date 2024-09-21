package persistence

import (
	"context"
	"errors"
	"log"

	"github.com/ei-sugimoto/tatekae/api/infrastructure"
	"github.com/ei-sugimoto/tatekae/api/infrastructure/ent"
	"github.com/ei-sugimoto/tatekae/api/infrastructure/ent/project"
	"github.com/ei-sugimoto/tatekae/api/infrastructure/ent/user"
	"github.com/ei-sugimoto/tatekae/api/model"
	"github.com/ei-sugimoto/tatekae/api/repo"
)

type PersistBill struct {
	db *infrastructure.DB
}

func NewPersistBill(db *infrastructure.DB) repo.BillRepo {
	return &PersistBill{
		db: db,
	}
}

func (p *PersistBill) Create(new *model.Bill) (*model.Bill, error) {
	ctx := context.Background()
	log.Println("new", new)
	dstUserID, err := p.db.User.Query().Where(user.ID(new.DstUser)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("dst user not found")
		}
		return nil, err
	}
	srcUserID, err := p.db.User.Query().Where(user.ID(new.SrcUser)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("src user not found")
		}
		return nil, err
	}

	ProjectID, err := p.db.Project.Query().Where(project.ID(new.ProjectID)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("project not found")
		}
		return nil, err
	}

	res, err := p.db.Bill.Create().SetDstUser(dstUserID).SetSrcUser(srcUserID).SetPrice(new.Price).SetProject(ProjectID).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &model.Bill{
		ID:        res.ID,
		Price:     res.Price,
		ProjectID: res.Edges.Project.ID,
		SrcUser:   res.Edges.SrcUser.ID,
		DstUser:   res.Edges.DstUser.ID,
	}, nil
}
