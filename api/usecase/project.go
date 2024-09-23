package usecase

import (
	"github.com/ei-sugimoto/tatekae/api/model"
	"github.com/ei-sugimoto/tatekae/api/repo"
)

type IProjectUsecase interface {
	Create(*model.Project) (*model.Project, error)
	List() ([]*model.Project, error)
	Join(int, int) error
	JoinList(int) ([]*model.User, error)
	Get(int) (*model.Project, error)
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

func (u *ProjectUsecase) Join(projectID, userID int) error {
	return u.repo.Join(projectID, userID)
}

func (u *ProjectUsecase) JoinList(projectID int) ([]*model.User, error) {
	return u.repo.JoinList(projectID)
}

func (u *ProjectUsecase) Get(projectID int) (*model.Project, error) {
	return u.repo.Get(projectID)
}
