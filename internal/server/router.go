package server

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/oxtx/go-rest-api/internal/handler"
	"github.com/oxtx/go-rest-api/internal/server/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type RouterParams struct {
	UserHandler *handler.UserHandler
	Logger      *slog.Logger
	Env         string
}

func NewRouter(p RouterParams) *gin.Engine {
	if p.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(
		middleware.RequestID(),
		middleware.Recovery(),
		middleware.Logging(p.Logger),
		middleware.CORS(),
	)

	r.GET("/healthz", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) })

	v1 := r.Group("/api/v1")
	{
		v1.POST("/users", p.UserHandler.Create)
		v1.GET("/users/:id", p.UserHandler.Get)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
