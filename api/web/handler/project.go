package handler

import (
	"context"
	"time"

	"github.com/ei-sugimoto/tatekae/api/model"
	"github.com/ei-sugimoto/tatekae/api/usecase"
	"github.com/ei-sugimoto/tatekae/api/web/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProjectHandler struct {
	ProjectUsecase usecase.IProjectUsecase
	proto.UnimplementedProjectServiceServer
}

func NewProjectHandler(projectUsecase usecase.IProjectUsecase) *ProjectHandler {
	return &ProjectHandler{
		ProjectUsecase: projectUsecase,
	}
}

func (h *ProjectHandler) Create(c context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
	MyID := c.Value("id").(int)
	new := &model.Project{
		Name:      req.Name,
		CreatedAt: time.Now(),
		CreatedBy: MyID,
	}

	res, err := h.ProjectUsecase.Create(new)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create project: %v", err)
	}

	return &proto.CreateResponse{Id: int32(res.ID), Name: res.Name}, nil
}
