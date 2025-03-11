package http

import (
	"bench_press_calculator/internal/http/responce"
	"bench_press_calculator/internal/lib/logger/sl"
	"bench_press_calculator/internal/model"
	"log/slog"
	"net/http"
	_ "bench_press_calculator/docs" 

	
	"github.com/go-chi/render"
)
type CreateRequest struct {
	Email    string  `json:"email" validate:"required,email"`
	Password string  `json:"password" validate:"required"`
	Weight   float32 `json:"weight"`
	Quantity float32 `json:"quantity"`
}
type CreateResponce struct {
	MaxWeight     float32 `json:"max weight: "`
	PercentBetter float32 `json:"better then average on: "`
}
// Create create new user
// @Summary create new user
// @Description create new user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body Request true "CreateRequest"
// @Success 201 {object} CreateRequest
// @Router /create [post]
func (s *Server) Create() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		const op = "transporet.http.Create"
		log := s.Logger.With(slog.String("op", op))

		var r CreateRequest
		err := render.DecodeJSON(req.Body, &r)
		if err != nil {
			log.Error("failed to decode request body", sl.Err(err))
			render.JSON(resp, req, responce.Error("failed to decode request"))
			return
		}

		log.Info("create", slog.Any("request", r))

		user := &model.User{
			Email:    r.Email,
			Password: r.Password,
		}

		err = user.BeforeCreate()
		if err != nil {
			log.Error("failed to prepare user", sl.Err(err))
			render.JSON(resp, req, responce.Error("failed to decode request"))
			return
		}

		log.Info("Creating user with data", slog.String("Email", user.Email), slog.String("Password", user.Password))

		stat, err := s.Svc.Calculate(user, r.Weight, r.Quantity)
		if err != nil {
			log.Error("failed to calculate stats", sl.Err(err))
			render.JSON(resp, req, responce.Error("failed to calculate stats"))
			return
		}

		render.JSON(resp, req, CreateResponce{
			MaxWeight:     stat.MaxPress,
			PercentBetter: stat.PersentBetter,
		})
	}
}

