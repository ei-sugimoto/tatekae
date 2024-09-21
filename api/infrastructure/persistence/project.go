package persistence

import (
	"context"

	"github.com/ei-sugimoto/tatekae/api/infrastructure"
	"github.com/ei-sugimoto/tatekae/api/model"
)

type PersistProject struct {
	db *infrastructure.DB
}

func NewPersistProject(db *infrastructure.DB) *PersistProject {
	return &PersistProject{
		db: db,
	}
}

func (p *PersistProject) Create(new *model.Project) (*model.Project, error) {
	ctx := context.Background()
	res, err := p.db.Project.Create().SetName(new.Name).SetCreatedAt(new.CreatedAt).SetCreatedBy(new.CreatedBy).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &model.Project{
		ID:        res.ID,
		Name:      res.Name,
		CreatedAt: res.CreatedAt,
		CreatedBy: res.CreatedBy,
	}, nil
}

func (p *PersistProject) List() ([]*model.Project, error) {
	ctx := context.Background()
	projects, err := p.db.Project.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]*model.Project, 0, len(projects))
	for _, project := range projects {
		res = append(res, &model.Project{
			ID:        project.ID,
			Name:      project.Name,
			CreatedAt: project.CreatedAt,
			CreatedBy: project.CreatedBy,
		})
	}

	return res, nil
}
