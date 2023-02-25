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
	"codeview/internal/persistence"

	imagehandler "codeview/internal/handler/http/image"
	problemhandler "codeview/internal/handler/http/problem"
	imagerepository "codeview/internal/repository/impl/image"
	problemrepository "codeview/internal/repository/impl/problem"
	imageservice "codeview/internal/service/impl/image"
	problemservice "codeview/internal/service/impl/problem"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	server *http.Server
}

func Init(cfg config.AppConfig) (*Server, error) {

	router := gin.Default()
	router.Use(cors.Default())

	persistence, err := persistence.Init(cfg)
	if err != nil {
		return nil, err
	}

	defer persistence.Close()

	// Initialize application repositories
	imageRepo := imagerepository.New(cfg, persistence.GCStorage)
	problemRepo := problemrepository.New(cfg, persistence.Postgres)

	// Initialize application services
	problemService := problemservice.New(cfg, problemRepo)
	imageService := imageservice.New(cfg, imageRepo)

	// Initialize application handler
	problemhandler.New(cfg, router, problemService)
	imagehandler.New(cfg, router, imageService)

	server := &http.Server{
		Addr:    cfg.RestServer.Port,
		Handler: router,
	}

	return &Server{
		server: server,
	}, nil

}

func (s *Server) Start() {
	// Graceful server shutdown - https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/server.go
	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to initialize server: %v\n", err)
		}
	}()

	log.Printf("Listening on port %v\n", s.server.Addr)

	// Wait for kill signal of channel
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// This blocks until a signal is passed into the quit channel
	<-quit

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown server
	log.Println("Shutting down server...")
	if err := s.server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}
}
