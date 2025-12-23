package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	psqlConnection "github.com/Chandra5468/todo-go/internal/database/postgresql"
	todoDeliveryHandler "github.com/Chandra5468/todo-go/internal/modules/todos/delivery"
	todosRepo "github.com/Chandra5468/todo-go/internal/modules/todos/repository"
	todosService "github.com/Chandra5468/todo-go/internal/modules/todos/service"
	usersDeliveryHandler "github.com/Chandra5468/todo-go/internal/modules/users/delivery"
	usersRepo "github.com/Chandra5468/todo-go/internal/modules/users/repository"
	usersService "github.com/Chandra5468/todo-go/internal/modules/users/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("unabel to load env file %v", err)
	}

	// 1. Creating a global logger using slog
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// 2. Database connection management
	psqlPool, err := psqlConnection.Connect(os.Getenv("GOOSE_DBSTRING"))
	if err != nil {
		log.Fatalf("unable to connect to postgres db pool %v", err)
	}
	// if application/http server is shutdown then need to close psql pool gracefully
	defer func() {
		slog.Warn("closing db pool")
		psqlPool.Close()
	}()

	// 3. Setting up service -> Handlers -> main. Also to handlers send routers
	// // // Add some chi custom middlewares
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	todosRepository := todosRepo.NewRepo(psqlPool)
	todosService := todosService.NewService(todosRepository)
	todoDeliveryHandler.NewHandler(r, todosService)

	userRepository := usersRepo.NewRepo(psqlPool)
	usersServices := usersService.NewService(userRepository)
	usersDeliveryHandler.NewHandler(r, usersServices)

	// 4. http Server declaration
	srv := &http.Server{
		Addr:         os.Getenv("http_addr"),
		Handler:      r,
		WriteTimeout: time.Second * 60,
		ReadTimeout:  time.Second * 60,
		IdleTimeout:  time.Minute,
	}

	// 5. Root context : will be cancelled on SIGINT/TERM
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	go func() {
		slog.Info("server started", "addr", os.Getenv("PORT"))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("server closed", "error", err)
			// stop()
			panic(err)
		}
	}()

	<-ctx.Done()
	slog.Info("shutdown signal received")

	// 6. Giving ongoing requests some time to finish. And closing db gracefully
	shutDownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutDownCtx); err != nil {
		slog.Error("server shutdown failed", "error", err)
	}

	slog.Info("graceful shutdown completed")

}
