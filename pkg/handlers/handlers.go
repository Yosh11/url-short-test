package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/Yosh11/url-short-test/database"
	"github.com/Yosh11/url-short-test/lib/genhash"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func initDB() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s\n", err.Error())
	}
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s\n", err.Error())
	}

	db, err := database.NewMSSQLDB(database.Config{
		NameDB:   viper.GetString("db.namedb"),
		Scheme:   viper.GetString("db.scheme"),
		User:     viper.GetString("db.user"),
		Password: os.Getenv("PASS_DB"), // private in .env
		Host:     viper.GetString("db.host"),
		Port:     viper.GetUint32("db.port"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s\n", err.Error())
	}
	return db
}

// AddURL set short url
func AddURL(ctx echo.Context) error {
	var req database.Urls
	db := initDB()

	err := ctx.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = ctx.Validate(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	ranHash := genhash.Generate()

	rURL := &database.Urls{
		Hash:  ranHash,
		URL:   req.URL,
		Count: 0,
	}

	db.Create(&rURL)

	return ctx.JSON(http.StatusOK, rURL)
}
