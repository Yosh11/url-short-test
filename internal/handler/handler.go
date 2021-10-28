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
	return nil
}