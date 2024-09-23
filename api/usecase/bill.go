package usecase

import (
	"github.com/ei-sugimoto/tatekae/api/model"
	"github.com/ei-sugimoto/tatekae/api/repo"
)

type IBillUsecase interface {
	Create(*model.Bill) (*model.Bill, error)
	ListByProject(int) ([]*model.Bill, error)
	Sumarize(int) ([]*model.SumarizeBill, error)
}

type BillUsecase struct {
	repo     repo.BillRepo
	userRepo repo.UserRepo
}

func NewBillUsecase(repo repo.BillRepo, userRepo repo.UserRepo) IBillUsecase {
	return &BillUsecase{
		repo:     repo,
		userRepo: userRepo,
	}
}

func (u *BillUsecase) Create(bill *model.Bill) (*model.Bill, error) {
	return u.repo.Create(bill)
}

func (u *BillUsecase) ListByProject(targetID int) ([]*model.Bill, error) {
	return u.repo.ListByProject(targetID)
}

func (u *BillUsecase) Sumarize(projectID int) ([]*model.SumarizeBill, error) {
	bills, err := u.repo.ListByProject(projectID)
	if err != nil {
		return nil, err
	}
	calcRes, err := model.CalcucateSummaries(bills)
	if err != nil {
		return nil, err
	}

	sumarize := make([]*model.SumarizeBill, 0, len(calcRes))
	SrcID := make([]int, 0, len(calcRes))
	DstID := make([]int, 0, len(calcRes))
	for _, v := range calcRes {
		SrcID = append(SrcID, v.SrcUser)
		DstID = append(DstID, v.DstUser)
	}

	Srcusers, err := u.userRepo.ListByIDs(SrcID)
	if err != nil {
		return nil, err
	}
	Dstusers, err := u.userRepo.ListByIDs(DstID)
	if err != nil {
		return nil, err
	}
	srcUserMap := make(map[int]*model.User)
	dstUserMap := make(map[int]*model.User)
	for _, user := range Srcusers {
		srcUserMap[user.ID] = user
	}
	for _, user := range Dstusers {
		dstUserMap[user.ID] = user
	}

	for _, v := range calcRes {
		sumarize = append(sumarize, &model.SumarizeBill{
			SrcUserName: srcUserMap[v.SrcUser].Username,
			DstUserName: dstUserMap[v.DstUser].Username,
			Amount:      v.Amount,
		})
	}
	return sumarize, nil
}
