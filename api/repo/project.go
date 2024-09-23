package repo

import "github.com/ei-sugimoto/tatekae/api/model"

type ProjectRepo interface {
	Create(*model.Project) (*model.Project, error)
	List() ([]*model.Project, error)
	Join(int, int) error
	Get(int) (*model.Project, error)
	JoinList(int) ([]*model.User, error)
}
