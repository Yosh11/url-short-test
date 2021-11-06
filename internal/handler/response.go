package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/Yosh11/url-short-test/init/err"
)

type errorResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, error error) {
	err.CheckError(error)

	c.AbortWithStatusJSON(statusCode, errorResponse{statusCode, error.Error()})
}
