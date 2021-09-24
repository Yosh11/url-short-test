package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Yosh11/url-short-test/internal/err"
	"github.com/Yosh11/url-short-test/internal/srv"
)

func main() {
	e := godotenv.Load()
	err.CheckFatal(e,"fail with env`s")
	err.InitLogrus()

	db := srv.InitMongo()
	s := new(srv.Server)
	go err.CheckError(s.Run(os.Getenv("PORT_API"), nil))

	graceful(s, db, 10*time.Second)
}

func graceful(s *srv.Server, db *mongo.Client, timeout time.Duration) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err.CheckError(db.Disconnect(ctx), "Database forced to shutdown")
	err.CheckFatal(s.Shutdown(ctx), "Server forced to shutdown")
	err.Info("Server exiting")
}