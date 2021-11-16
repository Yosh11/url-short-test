package main

import (
	"context"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/Yosh11/url-short-test/init/err"
	"github.com/Yosh11/url-short-test/init/srv"
	"github.com/Yosh11/url-short-test/internal/handler"
	"github.com/Yosh11/url-short-test/internal/repository"
	"github.com/Yosh11/url-short-test/internal/service"
)

func main() {
	// Init env`s
	e := godotenv.Load()
	err.CheckFatal(e,"fail with env`s")
	// Init custom logger
	err.InitLogrus()

	// Init repository (MongoDB)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := repository.InitMongo(ctx)
	repos := repository.NewRepository(ctx, db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	// New server
	s := new(srv.Server)
	go err.CheckError(s.Run(os.Getenv("PORT_API"), handlers.InitRoutes()))

	// Try shutdown app
	srv.GracefulShutdown(s, ctx, db, 10*time.Second)
}
