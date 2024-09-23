package handler

import (
	"context"
	"time"

	"connectrpc.com/connect"
	projectv1 "github.com/ei-sugimoto/tatekae/api/gen/proto_project/v1"
	"github.com/ei-sugimoto/tatekae/api/model"
	"github.com/ei-sugimoto/tatekae/api/usecase"
	"github.com/ei-sugimoto/tatekae/api/web/proto"
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

// func (h *ProjectHandler) Create(c context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
// 	MyID := c.Value("id").(int)
// 	new := &model.Project{
// 		Name:      req.Name,
// 		CreatedAt: time.Now(),
// 		CreatedBy: MyID,
// 	}

// 	res, err := h.ProjectUsecase.Create(new)
// 	if err != nil {
// 		return nil, status.Errorf(codes.Internal, "failed to create project: %v", err)
// 	}

// 	return &proto.CreateResponse{Id: int32(res.ID), Name: res.Name}, nil
// }

func (h *ProjectHandler) Create(ctx context.Context, arg *connect.Request[projectv1.ProjectServiceCreateRequest]) (*connect.Response[projectv1.ProjectServiceCreateResponse], error) {
	MyID := ctx.Value("id").(int)
	new := &model.Project{
		Name:      arg.Msg.Name,
		CreatedAt: time.Now(),
		CreatedBy: MyID,
	}

	res, err := h.ProjectUsecase.Create(new)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return &connect.Response[projectv1.ProjectServiceCreateResponse]{
		Msg: &projectv1.ProjectServiceCreateResponse{
			Id:   int32(res.ID),
			Name: res.Name,
		},
	}, nil
}

// func (h *ProjectHandler) List(c context.Context, req *proto.ListRequest) (*proto.ListResponse, error) {
// 	projects, err := h.ProjectUsecase.List()
// 	if err != nil {
// 		return nil, status.Errorf(codes.Internal, "failed to list projects: %v", err)
// 	}

// 	res := make([]*proto.Project, 0, len(projects))
// 	for _, project := range projects {
// 		res = append(res, &proto.Project{
// 			Id:   int32(project.ID),
// 			Name: project.Name,
// 		})
// 	}

// 	return &proto.ListResponse{Projects: res}, nil
// }

// func (h *ProjectHandler) Join(c context.Context, req *proto.JoinRequest) (*proto.JoinResponse, error) {
// 	MyID := c.Value("id").(int)
// 	MyUsername := c.Value("username").(string)
// 	err := h.ProjectUsecase.Join(int(req.Id), MyID)
// 	if err != nil {
// 		return nil, status.Errorf(codes.Internal, "failed to join project: %v", err)
// 	}

// 	return &proto.JoinResponse{Id: req.Id, Name: MyUsername}, nil
// }

func (h *ProjectHandler) List(ctx context.Context, arg *connect.Request[projectv1.ProjectServiceListRequest]) (*connect.Response[projectv1.ProjectServiceListResponse], error) {
	projects, err := h.ProjectUsecase.List()
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	res := make([]*projectv1.Project, 0, len(projects))
	for _, project := range projects {
		res = append(res, &projectv1.Project{
			Id:   int32(project.ID),
			Name: project.Name,
		})
	}

	return &connect.Response[projectv1.ProjectServiceListResponse]{
		Msg: &projectv1.ProjectServiceListResponse{
			Projects: res,
		},
	}, nil
}

func (h *ProjectHandler) Join(ctx context.Context, arg *connect.Request[projectv1.ProjectServiceJoinRequest]) (*connect.Response[projectv1.ProjectServiceJoinResponse], error) {
	myID := ctx.Value("id").(int)
	myUsername := ctx.Value("username").(string)
	id := int(arg.Msg.Id)

	err := h.ProjectUsecase.Join(id, myID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return &connect.Response[projectv1.ProjectServiceJoinResponse]{
		Msg: &projectv1.ProjectServiceJoinResponse{
			Id:   int32(id),
			Name: myUsername,
		},
	}, nil

}
