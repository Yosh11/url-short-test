package handlers

import (
	"log"
	"net/http"
	"sync"

	"github.com/Yosh11/url-short-test/database"
	"github.com/Yosh11/url-short-test/lib/genhash"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

type resBody struct {
	Long  string `json:"long"`
	Short string `json:"short"`
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

// InitDB init db
func InitDB() *gorm.DB {
	once.Do(func() {
		if err := initConfig(); err != nil {
			log.Panicf("error initializing configs: %s\n", err.Error())
		}
		var err error
		db, err = database.NewDB(database.Config{
			NameDB:   viper.GetString("db.namedb"),
			User:     viper.GetString("db.user"),
			Password: viper.GetString("db.pass"),
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
		})
		if err != nil {
			log.Panicf("failed to initialize db: %s\n", err.Error())
		}
	})
	return db
}

// AddURL set short url
func AddURL(ctx echo.Context) error {
	var req database.Urls
	db := InitDB()

	err := ctx.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = ctx.Validate(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	randHash := genhash.Generate()

	rURL := &database.Urls{
		Hash:   randHash,
		URL:    req.URL,
		Count:  0,
		Access: true,
		Code:   200,
	}

	db.Create(&rURL)

	return ctx.JSON(http.StatusOK, resBody{
		Long:  req.URL,
		Short: viper.GetString("srv.host") + viper.GetString("srv.port") + "/" + randHash,
	})
}

// RedirectURL redirect to long url
func RedirectURL(ctx echo.Context) error {
	var pattern database.Urls

	db := InitDB()

	hs := ctx.Param("hash")
	db.Where("hash = ?", hs).Find(&pattern)

	db.Model(&pattern).Update("count", pattern.Count+1)

	return ctx.Redirect(308, pattern.URL)
}

// GetInfo get information about short url
func GetInfo(ctx echo.Context) error {
	var pattern database.Urls

	db := InitDB()

	hs := ctx.Param("hash")
	db.Where("hash = ?", hs).Find(&pattern)

	return ctx.JSON(200, pattern)
}
