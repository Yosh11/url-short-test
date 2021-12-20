package srv

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Yosh11/url-short-test/init/log"
)

func GracefulShutdown(s *Server, ctx context.Context, db *mongo.Client, timeout time.Duration) {
	quit := make(chan os.Signal, 1)
	signal.Notify(
		quit,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGABRT)
	<-quit

	c, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	log.CheckError(db.Disconnect(c), "Database forced to shutdown")
	log.CheckFatal(s.Shutdown(c), "Server forced to shutdown")
	log.Info("Server exiting")
}
