package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Yosh11/url-short-test/database"
	"github.com/Yosh11/url-short-test/lib/validator"
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

	// Run and init server
	startServer()
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func startServer() {
	// Initialize Echo instance
	e := echo.New()
	e.Validator = validator.NewValidator()

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[INFO] method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())

	// Routes
	e.POST("/set", handlers.AddURL)

	s := &http.Server{
		Addr:           viper.GetString("srv.port"),
		Handler:        e,
		MaxHeaderBytes: 1 << 20, // 1 Mb
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	go func(s *http.Server) {
		if err := e.StartServer(s); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}(s)

	graceful(s, 5*time.Second)
}

// Graceful Shutdown implementation taken from echo doc
func graceful(s *http.Server, timeout time.Duration) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("error at graceful shutdown: %s", err.Error())
	} else {
		log.Println("Server stopped")
	}
}
