package usecase

import (
	"github.com/ei-sugimoto/tatekae/api/model"
	"github.com/ei-sugimoto/tatekae/api/repo"
)

type IBillUsecase interface {
	Create(*model.Bill) (*model.Bill, error)
}

type BillUsecase struct {
	repo repo.BillRepo
}

func NewBillUsecase(repo repo.BillRepo) IBillUsecase {
	return &BillUsecase{
		repo: repo,
	}
}

func (u *BillUsecase) Create(bill *model.Bill) (*model.Bill, error) {
	return u.repo.Create(bill)
}
