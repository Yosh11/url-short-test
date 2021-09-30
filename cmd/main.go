package main

import (
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/Yosh11/url-short-test/internal/err"
	"github.com/Yosh11/url-short-test/internal/srv"
)

func main() {
	// Init env`s
	e := godotenv.Load()
	err.CheckFatal(e,"fail with env`s")
	// Init custom logger
	err.InitLogrus()

	// Init repository (MondoDB)
	db := srv.InitMongo()

	// New server
	s := new(srv.Server)
	go err.CheckError(s.Run(os.Getenv("PORT_API"), nil))

	// Try shutdown app
	srv.GracefulShutdown(s, db, 10*time.Second)
}
