package repo

import "github.com/ei-sugimoto/tatekae/api/model"

type ProjectRepo interface {
	Create(*model.Project) (*model.Project, error)
}
