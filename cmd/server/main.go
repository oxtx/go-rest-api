package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/oxtx/go-rest-api/api/docs"
	"github.com/oxtx/go-rest-api/internal/config"
	"github.com/oxtx/go-rest-api/internal/handler"
	"github.com/oxtx/go-rest-api/internal/platform/logger"
	"github.com/oxtx/go-rest-api/internal/platform/postgres"
	"github.com/oxtx/go-rest-api/internal/repository"
	"github.com/oxtx/go-rest-api/internal/server"
	"github.com/oxtx/go-rest-api/internal/service"
)

// @title Go Rest API
// @version 1.0
// @description Example REST API using Gin
// @BasePath /

func main() {
	cfg := config.Load()
	logr := logger.New(cfg.Env)

	db, err := postgres.New(cfg.DBURL)
	if err != nil {
		log.Fatalf("db connect failed: %v", err)
	}

	userRepo := repository.NewUserRepository(db)
	userSvc := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userSvc)

	r := server.NewRouter(server.RouterParams{
		UserHandler: userHandler,
		Logger:      logr,
		Env:         cfg.Env,
	})

	srv := &http.Server{
		Addr:              ":" + cfg.HTTPPort,
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		logr.Info("server starting", "port", cfg.HTTPPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logr.Error("server error", "err", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logr.Error("shutdown error", "err", err)
	}
	logr.Info("server stopped")
}
