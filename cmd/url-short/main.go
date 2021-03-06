package main

import (
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    echoLog "github.com/labstack/gommon/log"
    "github.com/robfig/cron/v3"

    "github.com/Yosh11/url-short-test/lib/inspector"
    "github.com/Yosh11/url-short-test/lib/validator"
    "github.com/Yosh11/url-short-test/pkg/handlers"
)

func main() {
	// Run and init server
	startServer(":8080")
}

func startServer(addr string) {
	// Initialize Echo instance
	e := echo.New()
	// Initialize a new Cron job runner
	c := cron.New()
	e.Validator = validator.NewValidator()

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[INFO] method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
		LogLevel:  echoLog.ERROR,
	}))

	// Routes
	e.POST("/set", handlers.AddURL)
	e.GET("/:hash", handlers.RedirectURL)
	e.GET("/info/:hash", handlers.GetInfo)

	s := &http.Server{
		Addr:           addr,
		MaxHeaderBytes: 1 << 20, // 1 Mb
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	// Add and start job for cron
	c.AddFunc("@every 1m", inspector.Check)
	c.Start()

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
		log.Println("Server is graceful stopped")
	}
}
