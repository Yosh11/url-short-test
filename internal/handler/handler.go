package handler

import (
	"github.com/gin-gonic/gin"

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

	urls := router.Group("/urls")
	{
		urls.POST("/set", h.setNewUrl)
		urls.GET("/:hash", h.redirectUrl)
		urls.GET("/info/:hash", h.getInfoToUrl)
		urls.DELETE("/:hash", h.delUrl)
	}

	return router
}
