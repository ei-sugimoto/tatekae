package persistence

import (
	"context"

	"github.com/ei-sugimoto/tatekae/api/infrastructure"
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

	dstUserID := p.db.User.Query().Where(user.ID(new.DstUser)).OnlyX(ctx)
	srcUserID := p.db.User.Query().Where(user.ID(new.SrcUser)).OnlyX(ctx)
	ProjectID := p.db.Project.Query().Where(project.ID(new.ProjectID)).OnlyX(ctx)

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
