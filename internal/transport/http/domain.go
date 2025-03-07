package http

import (
	"bench_press_calculator/internal/config"
	"bench_press_calculator/internal/service"
	"log/slog"
	"net/http"
	_ "bench_press_calculator/docs"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	//swaggerFiles "github.com/swaggo/files"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type Server struct {
	server *http.Server
	router *chi.Mux
	logger *slog.Logger

	svc service.CalculatorService
}

func NewServer(cfg *config.Config, logger *slog.Logger, svc service.CalculatorService) *Server {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	return &Server{
		svc:    svc,
		router: router,
		logger: logger,
		server: &http.Server{
			Addr:        cfg.HTTPServer.Address,
			Handler:     router,
			ReadTimeout: cfg.HTTPServer.Timeout,
			IdleTimeout: cfg.HTTPServer.IdleTimeout,
		},
	}
}
func (s *Server) InitRoutes() {
	s.router.Route("/create", func(r chi.Router) {
		r.Use(middleware.AllowContentType("application/json"))
		r.Use(middleware.SetHeader("Content-Type", "application/json"))
		r.Post("/", s.Create())
	})
	s.router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

}
func (s *Server) Start() error {
	return s.server.ListenAndServe()
}
