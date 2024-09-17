package web

import (
	"os"

	"github.com/ei-sugimoto/tatekae/api/web/handler"
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
	v1 := r.Engine.Group("/v1")
	{
		v1.GET("/health", healthHandler.StatusOK)
	}
}
