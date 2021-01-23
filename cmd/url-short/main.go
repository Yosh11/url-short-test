package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Yosh11/url-short-test/database"
	"github.com/Yosh11/url-short-test/pkg/handlers"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func main() {
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
	if err = db.AutoMigrate(&database.Urls{}); err != nil {
		log.Fatalf("failed to mirgate: %s\n", err.Error())
	}
	if err := run(); err != nil {
		log.Fatalf("failed to start server: %s\n", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func run() error {
	// Initialize Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[INFO] method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())

	// Routes
	e.GET("/set", handlers.AddURL)

	s := &http.Server{
		Addr:           viper.GetString("srv.port"),
		MaxHeaderBytes: 1 << 20, // 1 Mb
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	e.Logger.Fatal(e.StartServer(s))

	return nil
}
