package http

import (
	responce "bench_press_calculator/internal/http/response"
	"bench_press_calculator/internal/lib/logger/sl"
	"bench_press_calculator/internal/model"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
)

func (s *Server) Create() http.HandlerFunc {

	type Request struct {
		Email    string  `json:"email" validate:"required,email"`
		Password string  `json:"password" validate:"required"`
		Weight   float32 `json:"weight"`
		Quantity float32 `json:"quantity"`
	}

	type Response struct {
		MaxWeight     float32 `json:"max weight: "`
		PercentBetter float32 `json:"better then average on: "`
	}

	return func(res http.ResponseWriter, req *http.Request) {
		const op = "handlers.creator.New"
		log := s.logger.With(slog.String("op", op))

		var r Request
		err := render.DecodeJSON(req.Body, &r)
		if err != nil {
			log.Error("failed to decode request body", sl.Err(err))
			render.JSON(res, req, responce.Error("failed to decode request"))
			return
		}

		log.Info("create", slog.Any("request", req))

		stat, err := s.svc.Calculate(&model.User{
			Email:    r.Email,
			Password: r.Password,
		},
			r.Weight,
			r.Quantity,
		)
		if err != nil {
			log.Error("failed to calculate stats", sl.Err(err))
			render.JSON(res, req, responce.Error("failed to calculate stats"))
			return
		}

		render.JSON(res, req, Response{
			MaxWeight:     stat.MaxPress,
			PercentBetter: stat.PercentBetter,
		})
	}
}
