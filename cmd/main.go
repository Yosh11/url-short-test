package main

import (
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/Yosh11/url-short-test/init/err"
	srv2 "github.com/Yosh11/url-short-test/init/srv"
	"github.com/Yosh11/url-short-test/internal/repository"
)

func main() {
	// Init env`s
	e := godotenv.Load()
	err.CheckFatal(e,"fail with env`s")
	// Init custom logger
	err.InitLogrus()

	// Init repository (MongoDB)
	db := repository.InitMongo()

	// New server
	s := new(srv2.Server)
	go err.CheckError(s.Run(os.Getenv("PORT_API"), nil))

	// Try shutdown app
	srv2.GracefulShutdown(s, db, 10*time.Second)
}
