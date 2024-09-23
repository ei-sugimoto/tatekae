package handler

import (
	"context"
	"time"

	"connectrpc.com/connect"
	userv1 "github.com/ei-sugimoto/tatekae/api/gen/proto_user/v1"
	"github.com/ei-sugimoto/tatekae/api/model"
	"github.com/ei-sugimoto/tatekae/api/pkg"
	"github.com/ei-sugimoto/tatekae/api/usecase"
	"github.com/ei-sugimoto/tatekae/api/web/proto"
)

type UserHandler struct {
	UserUsecase usecase.IUserUsecase

	proto.UnimplementedUserServiceServer
}

func NewUserHandler(userusecase usecase.IUserUsecase) *UserHandler {
	return &UserHandler{
		UserUsecase: userusecase,
	}
}

func (h *UserHandler) Register(ctx context.Context, arg *connect.Request[userv1.RegisterRequest]) (*connect.Response[userv1.RegisterResponse], error) {
	newUser := &model.User{
		Username: arg.Msg.Username,
		Email:    arg.Msg.Email,
		Password: arg.Msg.Password,
		CreateAt: time.Now(),
	}

	res, err := h.UserUsecase.Create(newUser)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	token, err := pkg.NewToken(res)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return &connect.Response[userv1.RegisterResponse]{
		Msg: &userv1.RegisterResponse{
			Token: token,
		},
	}, nil
}

func (h *UserHandler) Login(ctx context.Context, arg *connect.Request[userv1.LoginRequest]) (*connect.Response[userv1.LoginResponse], error) {
	user, err := h.UserUsecase.Login(arg.Msg.Email, arg.Msg.Password)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err.Err)
	}

	token, TokenErr := pkg.NewToken(user)
	if TokenErr != nil {
		return nil, connect.NewError(connect.CodeInternal, TokenErr)
	}

	return &connect.Response[userv1.LoginResponse]{
		Msg: &userv1.LoginResponse{
			Token: token,
		},
	}, nil
}

func (h *UserHandler) GetByID(ctx context.Context, arg *connect.Request[userv1.GetByIDRequest]) (*connect.Response[userv1.GetByIDResponse], error) {
	user, err := h.UserUsecase.Get(int(arg.Msg.Id))
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return &connect.Response[userv1.GetByIDResponse]{
		Msg: &userv1.GetByIDResponse{
			Id:       int32(user.ID),
			Username: user.Username,
			Email:    user.Email,
		},
	}, nil
}
