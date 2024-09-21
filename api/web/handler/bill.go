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
