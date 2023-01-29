package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"codeview/config"

	problemHandler "codeview/internal/handler/http/problem"
	imageRepository "codeview/internal/repository/image"
	problemService "codeview/internal/service/problem"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting server...")

	config := config.Get()

	router := gin.Default()

	// TODO: Separate this to a proper definition
	// Initialize google storage client
	log.Printf("Connecting to Cloud Storage\n")
	storage, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("error creating cloud storage client: %v", err)
	}

	// Initialize application dependencies
	imageRepo := imageRepository.New(storage, config.GC_BUCKET_NAME)
	problemService := problemService.New(imageRepo)
	problemHandler.New(router, problemService, &problemHandler.Config{
		MaxBodyBytes: config.MAX_BODY_BYTES,
	})

	srv := &http.Server{
		Addr:    config.REST_PORT,
		Handler: router,
	}

	// Graceful server shutdown - https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/server.go
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to initialize server: %v\n", err)
		}
	}()

	log.Printf("Listening on port %v\n", srv.Addr)

	// Wait for kill signal of channel
	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// This blocks until a signal is passed into the quit channel
	<-quit

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// shutdown google storage data source
	if err := storage.Close(); err != nil {
		log.Fatalf("A problem occurred gracefully shutting down data sources - error closing Cloud Storage client: %v\n", err)
	}

	// Shutdown server
	log.Println("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}
}
