package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/Yosh11/url-short-test/docs"
	"github.com/Yosh11/url-short-test/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{service: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	docs.SwaggerInfo.Title = "URL Shortener (the task)"
	docs.SwaggerInfo.Description = "This is a simple implementation of a shortening link service."
	docs.SwaggerInfo.Version = "2.1.1"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}

	// use ginSwagger middleware to serve the API docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/:hash", h.redirectUrl)

	urls := router.Group("/urls")
	{
		urls.POST("/set", h.setNewUrl)
		urls.GET("/info/:hash", h.getInfoToUrl)
		urls.DELETE("/:hash", h.delUrl)
	}

	return router
}
