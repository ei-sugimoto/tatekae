package handler

import (
	"context"

	"github.com/ei-sugimoto/tatekae/api/model"
	"github.com/ei-sugimoto/tatekae/api/usecase"
	"github.com/ei-sugimoto/tatekae/api/web/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BillHandler struct {
	u usecase.IBillUsecase
	proto.UnimplementedBillServiceServer
}

func NewBillHandler(u usecase.IBillUsecase) *BillHandler {
	return &BillHandler{
		u: u,
	}
}

func (h *BillHandler) Create(c context.Context, req *proto.CreateBillRequest) (*proto.CreateBillResponse, error) {
	new := &model.Bill{
		Price:     int(req.Price),
		ProjectID: int(req.ProjectId),
		SrcUser:   int(req.SrcUserId),
		DstUser:   int(req.DstUserId),
	}

	res, err := h.u.Create(new)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create project: %v", err)
	}

	return &proto.CreateBillResponse{
		ID:        int32(res.ID),
		Price:     int32(res.Price),
		ProjectId: int32(res.ProjectID),
		SrcUserId: int32(res.SrcUser),
		DstUserId: int32(res.DstUser),
	}, nil
}

func (h *BillHandler) ListByProject(c context.Context, req *proto.ListByProjectRequest) (*proto.ListByProjectResponse, error) {
	res, err := h.u.ListByProject(int(req.ProjectId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list project: %v", err)
	}

	bills := make([]*proto.Bill, 0, len(res))
	for _, v := range res {
		bills = append(bills, &proto.Bill{
			ID:        int32(v.ID),
			Price:     int32(v.Price),
			ProjectId: int32(v.ProjectID),
			SrcUserId: int32(v.SrcUser),
			DstUserId: int32(v.DstUser),
		})
	}

	return &proto.ListByProjectResponse{
		Bills: bills,
	}, nil
}

func (h *BillHandler) SumrizeByProject(c context.Context, req *proto.SumrizeRequest) (*proto.SumrizeResponse, error) {
	res, err := h.u.Sumarize(int(req.ProjectId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to sumarize project: %v", err)
	}

	sumarize := make([]*proto.SumrizeBill, 0, len(res))
	for _, v := range res {
		sumarize = append(sumarize, &proto.SumrizeBill{
			SrcUserName: v.SrcUserName,
			DstUserName: v.DstUserName,
			Price:       int32(v.Amount),
		})
	}

	return &proto.SumrizeResponse{
		SumrizeBills: sumarize,
	}, nil
}
