package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// AddURL set short url
func AddURL(c echo.Context) error {
	return c.String(http.StatusOK, "Ok")
}
