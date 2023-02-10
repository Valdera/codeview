package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"codeview/config"
	"codeview/persistence"

	imageHandler "codeview/internal/handler/http/image"
	problemHandler "codeview/internal/handler/http/problem"
	imageRepository "codeview/internal/repository/image"
	problemRepository "codeview/internal/repository/problem"
	imageService "codeview/internal/service/image"
	problemService "codeview/internal/service/problem"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start() {
	log.Println("Starting server...")

	config := config.Get()

	router := gin.Default()
	router.Use(cors.Default())

	db, err := persistence.Init(config)
	if err != nil {
		return
	}

	// Initialize application repositories
	imageRepo := imageRepository.New(db.GCStorage)
	problemRepo := problemRepository.New(db.Postgres)

	// Initialize application services
	problemService := problemService.New(problemRepo)
	imageService := imageService.New(imageRepo)

	// Initialize application handler
	problemHandler.New(router, problemService, &problemHandler.Config{})
	imageHandler.New(router, imageService, &imageHandler.Config{
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
	if err := db.Close(); err != nil {
		log.Fatalf("A problem occurred gracefully shutting down data sources - error closing Cloud Storage client: %v\n", err)
	}

	// Shutdown server
	log.Println("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}
}
