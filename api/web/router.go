package web

import (
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/ei-sugimoto/tatekae/api/infrastructure"
	"github.com/ei-sugimoto/tatekae/api/infrastructure/persistence"
	"github.com/ei-sugimoto/tatekae/api/usecase"
	"github.com/ei-sugimoto/tatekae/api/web/handler"
	"github.com/ei-sugimoto/tatekae/api/web/middleware"
	"github.com/ei-sugimoto/tatekae/api/web/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Router struct {
	Engine *grpc.Server
	Port   string
}

func NewRouter() *Router {
	engine := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			middleware.UnaryLoggingInterceptor(),
			middleware.AuthUnaryInterceptor,
		),
	)
	return &Router{Engine: engine}
}

func (r *Router) Run() {

	db := infrastructure.NewDB()

	db.Migrate()

	if port := os.Getenv("PORT"); port == "" {
		r.Port = ":8080"
	} else {
		r.Port = ":" + port
	}
	listener, err := net.Listen("tcp", r.Port)
	if err != nil {
		panic(err)
	}

	r.NewUserService()
	r.NewProjectService()

	reflection.Register(r.Engine)

	go func() {

		log.Println("Server is running on port", r.Port)
		if err := r.Engine.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}

	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")
	r.Engine.GracefulStop()

}

func (r *Router) NewUserService() {

	db := infrastructure.NewDB()
	userPersistence := persistence.NewPersistUser(db)
	userUsecase := usecase.NewUserUsecase(userPersistence)
	userHandler := handler.NewUserHandler(userUsecase)

	proto.RegisterUserServiceServer(r.Engine, userHandler)

}

func (r *Router) NewProjectService() {

	db := infrastructure.NewDB()
	projectPersistence := persistence.NewPersistProject(db)
	projectUsecase := usecase.NewProjectUsecase(projectPersistence)
	projectHandler := handler.NewProjectHandler(projectUsecase)

	proto.RegisterProjectServiceServer(r.Engine, projectHandler)

}
