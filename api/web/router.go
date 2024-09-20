package web

import (
	"os"

	"github.com/ei-sugimoto/tatekae/api/infrastructure"
	"github.com/ei-sugimoto/tatekae/api/infrastructure/persistence"
	"github.com/ei-sugimoto/tatekae/api/usecase"
	"github.com/ei-sugimoto/tatekae/api/web/handler"
	"github.com/ei-sugimoto/tatekae/api/web/middleware"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine *gin.Engine
	Port   string
}

func NewRouter() *Router {
	engine := gin.Default()
	return &Router{Engine: engine}
}

func (r *Router) Run() {
	if port := os.Getenv("PORT"); port == "" {
		r.Port = ":8080"
	} else {
		r.Port = ":" + port
	}
	r.SetRouting()
	r.Engine.Run(r.Port)
}

func (r *Router) SetRouting() {
	healthHandler := handler.NewHealthHandler()
	db := infrastructure.NewDB()
	userPersistence := persistence.NewPersistUser(db)
	userUsecase := usecase.NewUserUsecase(userPersistence)
	userHandler := handler.NewUserHandler(userUsecase)

	db.Migrate()

	v1 := r.Engine.Group("/v1")
	{
		v1.GET("/health", healthHandler.StatusOK)
		v1.POST("/user/register", userHandler.Register)
		v1.POST("/user/login", userHandler.Login)
	}
	authed := r.Engine.Group("/v1")
	authed.Use(middleware.Auth)
	{
	}
}
