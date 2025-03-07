package main

import (
	_ "bench_press_calculator/docs"

	"bench_press_calculator/internal/config"
	"bench_press_calculator/internal/lib/logger/sl"
	"bench_press_calculator/internal/service/calculator"
	"bench_press_calculator/internal/storage/postgresql"
	server "bench_press_calculator/internal/transport/http"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envProd  = "prod"
	envDev   = "dev"
)

// @title Documenting API (Wehw93)
// @version 1
// @Description Sample description

// @contact.name Egor Titov
// @contact.url https://github.com/wehw93
// @contact.email wehw93@mail.ru

// @host localhost:8080

func main() {
	cfg := config.MustLoad()

	log := SetupLogger(cfg.Env)
	log.Info("starting bench press calculator")

	store, err := postgresql.New(cfg.DB.GetDSN())
	if err != nil {
		panic(err)
	}
	defer store.Close()

	svc := calculator.NewService(store)

	srv := server.NewServer(&cfg, log, svc)
	srv.InitRoutes()

	log.Info("starting server", slog.String("addr", cfg.HTTPServer.Address))

	if err := srv.Start(); err != nil {
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
