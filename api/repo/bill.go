package repo

import "github.com/ei-sugimoto/tatekae/api/model"

type BillRepo interface {
	Create(new *model.Bill) (*model.Bill, error)
}
