package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Yosh11/url-short-test/internal/models"
)

func (h *Handler) setNewUrl(c *gin.Context) {
	var input models.SetUrl
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}
	// TODO check to valid url (spell check and response status on request)
	res, err := h.service.SetUrl(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) redirectUrl(c *gin.Context) {
	hash := c.Param("hash")
	if hash == "" {
		newErrorResponse(c, http.StatusBadRequest, errors.New("don`t have hash key"))
		return
	}

	res, err := h.service.GetUrl(hash)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.Redirect(http.StatusPermanentRedirect, res)
}


func (h *Handler) getInfoToUrl(c *gin.Context) {
	hash := c.Param("hash")
	if hash == "" {
		newErrorResponse(c, http.StatusBadRequest, errors.New("don`t have hash key"))
		return
	}

	res, err := h.service.GetUrlInfo(hash)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) delUrl(c *gin.Context) {
	hash := c.Param("hash")
	if hash == "" {
		newErrorResponse(c, http.StatusBadRequest, errors.New("don`t have hash key"))
		return
	}

	err := h.service.DeleteUrl(hash)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "url removed",
	})
}
