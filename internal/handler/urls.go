package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Yosh11/url-short-test/internal/models"
)

// setNewUrl godoc
// @Summary      Set new url
// @Description  Accepts a link and returns a shortened version which will redirect to the original source.
// @Accept       json
// @Produce      json
// @Param        url body models.SetUrl true "Your url"
// @Tags         /urls
// @Success      200 {object} models.SetUrlResp
// @Failure      400 {object} handler.errorResponse
// @Router       /urls/set [post]
func (h *Handler) setNewUrl(c *gin.Context) {
	var input models.SetUrl
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	httpRes, err := http.Get(input.Url)
	if err != nil || httpRes.StatusCode != 200 {
		newErrorResponse(c, http.StatusBadRequest,
			errors.New("the URL passed is not valid. the site is not available"))
		return
	}

	res, err := h.service.SetUrl(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// redirectUrl godoc
// @Summary      Redirect to site
// @Description  Redirect to your site using a short link
// @Param        hash path string true "Uniq hash"
// @Tags         /
// @Success      301
// @Router       /{hash} [get]
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

	c.Redirect(http.StatusPermanentRedirect, res.Url)
}

// getInfoToUrl godoc
// @Summary      Get information of your url
// @Description  Takes a hash of your post in a link and displays information about your URL
// @Param        hash path string true "Uniq hash"
// @Tags         /urls
// @Success      200 {object} models.Url
// @Failure      400 {object} handler.errorResponse
// @Router       /urls/info/{hash} [get]
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

// setNewUrl godoc
// @Summary      Delete your url from service
// @Description  Takes the hash of your entry in the link and logically deletes the entry
// @Param        hash path string true "Uniq hash"
// @Tags         /urls
// @Success      200
// @Failure      400 {object} handler.errorResponse
// @Router       /urls/{hash} [delete]
func (h *Handler) delUrl(c *gin.Context) {
	hash := c.Param("hash")
	if hash == "" {
		newErrorResponse(c, http.StatusBadRequest, errors.New("don`t have hash key"))
		return
	}

	obj, err := h.service.GetUrlInfo(hash)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err)
		return
	}

	err = h.service.DeleteUrl(obj.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "url removed",
	})
}
