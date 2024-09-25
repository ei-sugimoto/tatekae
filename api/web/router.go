package web

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	"github.com/ei-sugimoto/tatekae/api/gen/proto_bill/v1/billv1connect"
	"github.com/ei-sugimoto/tatekae/api/gen/proto_project/v1/projectv1connect"
	"github.com/ei-sugimoto/tatekae/api/gen/proto_user/v1/userv1connect"
	"github.com/ei-sugimoto/tatekae/api/infrastructure"
	"github.com/ei-sugimoto/tatekae/api/infrastructure/persistence"
	"github.com/ei-sugimoto/tatekae/api/usecase"
	"github.com/ei-sugimoto/tatekae/api/web/handler"
	"github.com/ei-sugimoto/tatekae/api/web/middleware"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Router struct {
	Engine *http.ServeMux
	Port   string
}

func NewRouter() *Router {
	engine := http.NewServeMux()
	reflrector := grpcreflect.NewStaticReflector(
		"proto_bill.v1.BillService",
		"proto_project.v1.ProjectService",
		"proto_user.v1.UserService",
	)

	engine.Handle(grpcreflect.NewHandlerV1(reflrector))
	engine.Handle(grpcreflect.NewHandlerV1Alpha(reflrector))

	return &Router{Engine: engine}
}

func (r *Router) Run() {

	db := infrastructure.NewDB()

	db.Migrate()

	r.SetHandler()

	server := &http.Server{
		Addr:    ":8080",
		Handler: cors.AllowAll().Handler(h2c.NewHandler(r.Engine, &http2.Server{})),
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	db.Close()

}

func (r *Router) SetHandler() {
	db := infrastructure.NewDB()

	UserPer := persistence.NewPersistUser(db)
	ProjectPer := persistence.NewPersistProject(db)
	BillPer := persistence.NewPersistBill(db)

	UserUsecase := usecase.NewUserUsecase(UserPer)
	ProjectUsecase := usecase.NewProjectUsecase(ProjectPer)
	BillUsecase := usecase.NewBillUsecase(BillPer, UserPer)

	UserHandler := handler.NewUserHandler(UserUsecase)
	ProjectHandler := handler.NewProjectHandler(ProjectUsecase)
	BillHandler := handler.NewBillHandler(BillUsecase)

	authInterceptor := connect.WithInterceptors(middleware.AuthUnaryInterceptor())
	logInterceptor := connect.WithInterceptors(middleware.LoggingInterceptor())

	r.Engine.Handle(userv1connect.NewUserServiceHandler(UserHandler, logInterceptor))
	r.Engine.Handle(projectv1connect.NewProjectServiceHandler(ProjectHandler, logInterceptor, authInterceptor))
	r.Engine.Handle(billv1connect.NewBillServiceHandler(BillHandler, logInterceptor, authInterceptor))

}
