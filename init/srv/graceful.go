package srv

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Yosh11/url-short-test/init/err"
)

func GracefulShutdown(s *Server, ctx context.Context, db *mongo.Client, timeout time.Duration) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	c, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	err.CheckError(db.Disconnect(c), "Database forced to shutdown")
	err.CheckFatal(s.Shutdown(c), "Server forced to shutdown")
	err.Info("Server exiting")
}