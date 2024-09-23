package persistence

import (
	"context"

	"github.com/ei-sugimoto/tatekae/api/infrastructure"
	"github.com/ei-sugimoto/tatekae/api/infrastructure/ent/project"
	"github.com/ei-sugimoto/tatekae/api/model"
	"github.com/ei-sugimoto/tatekae/api/repo"
)

type PersistProject struct {
	db *infrastructure.DB
}

func NewPersistProject(db *infrastructure.DB) repo.ProjectRepo {
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

func (p *PersistProject) Join(projectID, userID int) error {
	ctx := context.Background()
	_, err := p.db.Project.UpdateOneID(projectID).AddUserIDs(userID).Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (p *PersistProject) JoinList(projectID int) ([]*model.User, error) {
	ctx := context.Background()
	project, err := p.db.Project.Query().Where(project.ID(projectID)).WithUsers().Only(ctx)
	if err != nil {
		return nil, err
	}

	users := project.Edges.Users
	res := make([]*model.User, 0, len(users))
	for _, user := range users {
		res = append(res, &model.User{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		})
	}

	return res, nil
}

func (p *PersistProject) Get(projectID int) (*model.Project, error) {
	ctx := context.Background()
	project, err := p.db.Project.Query().Where(project.ID(projectID)).Only(ctx)
	if err != nil {
		return nil, err
	}

	return &model.Project{
		ID:        project.ID,
		Name:      project.Name,
		CreatedAt: project.CreatedAt,
		CreatedBy: project.CreatedBy,
	}, nil
}
