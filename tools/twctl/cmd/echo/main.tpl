package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"


	{{.importPackages}}
)

var configFile = flag.String("f", "etc/{{.serviceName}}.yaml", "the config file")

func main() {
	flag.Parse()

	// Load the configuration file
	var c config.Config
	conf.MustLoad(*configFile, &c)

	// Create a new Echo instance
	e := echo.New()
	e.Use(middleware.Recover()) // Recovery middleware

	// Create a new service context
	svcCtx := svc.NewServiceContext(c)
	handler.RegisterHandlers(e, svcCtx)

	// Create a context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Start server
	go func() {
		log.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
		if err := e.Start(fmt.Sprintf("%s:%d", c.Host, c.Port)); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Shutting down the server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	<-ctx.Done()
	log.Println("Received shutdown signal, shutting down gracefully...")

	// Create a context with a timeout of 1 second.
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := e.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Server shutdown complete")
}