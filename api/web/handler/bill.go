package handler

import (
	"context"

	"connectrpc.com/connect"
	billv1 "github.com/ei-sugimoto/tatekae/api/gen/proto_bill/v1"
	"github.com/ei-sugimoto/tatekae/api/model"
	"github.com/ei-sugimoto/tatekae/api/usecase"
	"github.com/ei-sugimoto/tatekae/api/web/proto"
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

func (h *BillHandler) Create(ctx context.Context, arg *connect.Request[billv1.BillServiceCreateRequest]) (*connect.Response[billv1.BillServiceCreateResponse], error) {
	new := &model.Bill{
		Price:     int(arg.Msg.Price),
		ProjectID: int(arg.Msg.ProjectId),
		SrcUser:   int(arg.Msg.SrcUserId),
		DstUser:   int(arg.Msg.DstUserId),
	}

	res, err := h.u.Create(new)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&billv1.BillServiceCreateResponse{
		Id:        int32(res.ID),
		Price:     int32(res.Price),
		ProjectId: int32(res.ProjectID),
		SrcUserId: int32(res.SrcUser),
		DstUserId: int32(res.DstUser),
	}), nil
}

func (h *BillHandler) SumarizeByProject(ctx context.Context, arg *connect.Request[billv1.BillServiceSumarizeByProjectRequest]) (*connect.Response[billv1.BillServiceSumarizeByProjectResponse], error) {
	res, err := h.u.Sumarize(int(arg.Msg.ProjectId))
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	sumarize := make([]*billv1.SumrizeBill, 0, len(res))
	for _, v := range res {
		sumarize = append(sumarize, &billv1.SumrizeBill{
			SrcUserName: v.SrcUserName,
			DstUserName: v.DstUserName,
			Price:       int32(v.Amount),
		})
	}

	return connect.NewResponse(&billv1.BillServiceSumarizeByProjectResponse{
		Bills: sumarize,
	}), nil

}

func (h *BillHandler) List(ctx context.Context, arg *connect.Request[billv1.BillServiceListRequest]) (*connect.Response[billv1.BillServiceListResponse], error) {
	res, err := h.u.ListByProject(int(arg.Msg.ProjectId))
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	bills := make([]*billv1.Bill, 0, len(res))
	for _, v := range res {
		bills = append(bills, &billv1.Bill{
			Id:          int32(v.ID),
			Price:       int32(v.Price),
			SrcUserName: v.SrcName,
			DstUserName: v.DstName,
		})
	}

	return connect.NewResponse(&billv1.BillServiceListResponse{
		Bills: bills,
	}), nil
}
