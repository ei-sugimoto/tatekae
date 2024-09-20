package handler

import (
	"context"
	"time"

	"github.com/ei-sugimoto/tatekae/api/model"
	"github.com/ei-sugimoto/tatekae/api/pkg"
	"github.com/ei-sugimoto/tatekae/api/usecase"
	proto "github.com/ei-sugimoto/tatekae/api/web/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (h *UserHandler) Register(c context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	newUser := &model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		CreateAt: time.Now(),
	}

	res, err := h.UserUsecase.Create(newUser)
	if err != nil {

		return nil, status.Errorf(codes.Internal, "failed to create user: %v", err)
	}

	token, err := pkg.NewToken(res)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate token: %v", err)
	}

	return &proto.RegisterResponse{Token: token}, nil

}

func (h *UserHandler) Login(c context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	user, err := h.UserUsecase.Login(req.Email, req.Password)
	if err.Err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user: %v", err.Err)
	}

	token, TokenErr := pkg.NewToken(user)
	if TokenErr != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user: %v", TokenErr)
	}

	return &proto.LoginResponse{Token: token}, nil
}
