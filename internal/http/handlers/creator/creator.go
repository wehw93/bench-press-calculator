package creator

import (
	resp "bench_press_calculator/internal/http/responce"
	"bench_press_calculator/internal/lib/logger/sl"
	"bench_press_calculator/internal/model"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
)

type HTTPCreator interface {
	CreateUser(u *model.User) error
	Calculate(user *model.User, weight float32, quantity float32) (*model.Stat, error)
}

func New(log *slog.Logger, httpCreator HTTPCreator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.creator.New"

		log := log.With(
			slog.String("op", op),
		)

		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if errors.Is(err, io.EOF) {
			log.Error("request body is empty")

			render.JSON(w, r, resp.Error("empty request"))
			return
		}
		if err != nil {
			log.Error("failed to decode requset body", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to decode request"))
			return
		}
		log.Info("request body decoded", slog.Any("Requset", err))
		user := &model.User{
			Email:    req.Email,
			Password: req.Password,
			Weight:   req.Weight,
		}
		err = user.BeforeCreate()
		if err != nil {
			log.Error("failed to prepare user", sl.Err(err))
			render.JSON(w, r, resp.Error("invalid user data"))
			return
		}
		err = httpCreator.CreateUser(user)
		if err != nil {
			log.Error("failed to create user", sl.Err(err))
			render.JSON(w, r, resp.Error("failed to create user"))
			return
		}
		stat, err := httpCreator.Calculate(user, req.Weight, req.Quantity)
		if err != nil {
			log.Error("failed to calculate stats", sl.Err(err))
			render.JSON(w, r, resp.Error("failed to calculate stats"))
			return
		}
		response := &Responce{
			MaxWeight:     stat.MaxPress,
			PercentBetter: 0,
		}
		render.JSON(w, r, resp.OK())
		render.JSON(w, r, response)

	}
}
