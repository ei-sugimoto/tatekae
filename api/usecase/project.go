package usecase

import (
	"github.com/ei-sugimoto/tatekae/api/model"
	"github.com/ei-sugimoto/tatekae/api/repo"
)

type IProjectUsecase interface {
	Create(*model.Project) (*model.Project, error)
	List() ([]*model.Project, error)
}

type ProjectUsecase struct {
	repo repo.ProjectRepo
}

func NewProjectUsecase(repo repo.ProjectRepo) IProjectUsecase {
	return &ProjectUsecase{
		repo: repo,
	}
}

func (u *ProjectUsecase) Create(project *model.Project) (*model.Project, error) {
	return u.repo.Create(project)
}

func (u *ProjectUsecase) List() ([]*model.Project, error) {
	return u.repo.List()
}
