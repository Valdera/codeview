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
	handler "codeview/internal/handler/http"
	"codeview/internal/persistence"
	repository "codeview/internal/repository/impl"
	service "codeview/internal/service/impl"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Server struct {
	server      *http.Server
	persistence *persistence.Persistence
}

func Init(cfg config.AppConfig) (*Server, error) {

	router := gin.Default()
	router.Use(cors.Default())

	persistence, err := persistence.Init(cfg)
	if err != nil {
		return nil, err
	}

	router.Use(sessions.Sessions("sessions", persistence.SessionStore))

	// Initialize application repositories
	imageRepo := repository.NewImageRepository(cfg, persistence.GCStorage)
	problemRepo := repository.NewProblemRepository(cfg, persistence.Postgres)
	userRepo := repository.NewUserRepository(cfg, persistence.Postgres)

	// Initialize application services
	problemService := service.NewProblemService(cfg, problemRepo)
	imageService := service.NewImageService(cfg, imageRepo)
	authservice := service.NewAuthService(cfg, userRepo)

	// Initialize application handler
	handler.InitProblemHandler(cfg, router, problemService)
	handler.InitImageHandler(cfg, router, imageService)
	handler.InitAuthHandler(cfg, router, authservice)

	server := &http.Server{
		Addr:    cfg.RestServer.Port,
		Handler: router,
	}

	return &Server{
		server:      server,
		persistence: persistence,
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

	// Close persistence connection
	if err := s.persistence.Close(); err != nil {
		log.Fatalf("A problem occurred gracefully shutting down data sources: %v\n", err)
	}

	// Shutdown server
	log.Println("Shutting down server...")
	if err := s.server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}
}
