package handlers

import (
	"net/http"

	"github.com/Yosh11/url-short-test/database"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

// AddURL set short url
func AddURL(ctx echo.Context) error {
	var req database.Urls
	err := ctx.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = ctx.Validate(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	return ctx.JSON(http.StatusOK, req.URL)
}
