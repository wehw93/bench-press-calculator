package main

import (
	"bench_press_calculator/internal/config"
	"bench_press_calculator/internal/http/handlers/creator"
	"bench_press_calculator/internal/lib/logger/sl"
	"bench_press_calculator/internal/storage/postgresql"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

const (
	envLocal = "local"
	envProd  = "prod"
	envDev   = "dev"
)

func main() {
	cfg := config.MustLoad()

	log := SetupLogger(cfg.Env)
	log.Info("starting press branch calculator")

	store, err := postgresql.New(cfg.СonnString)
	if err != nil {
		panic(err)
	}
	defer store.Close()

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	router.Route("/create", func(r chi.Router) {
		r.Use(middleware.AllowContentType("application/json"))
		r.Use(middleware.SetHeader("Content-Type", "application/json"))
		r.Post("/", creator.New(log, store.User()))
	})
	srv := &http.Server{
		Addr:        cfg.Addres,
		Handler:     router,
		ReadTimeout: cfg.Timeout,
		IdleTimeout: cfg.Idle_timeout,
	}
	log.Info("starting server", slog.String("addr", srv.Addr))
	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server", sl.Err(err))
		os.Exit(1)

	}
}

func SetupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}
