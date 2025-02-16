package creator

import (
	resp "bench_press_calculator/internal/http/responce"
	"bench_press_calculator/internal/lib/logger/sl"
	"bench_press_calculator/internal/model"
	"bytes"
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
		log := log.With(slog.String("op", op))

		// Читаем тело запроса и логируем его
		body, _ := io.ReadAll(r.Body)
		log.Info("Raw request body:", slog.String("body", string(body)))
		r.Body = io.NopCloser(bytes.NewReader(body)) // Восстанавливаем r.Body

		var req Request
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode request body", sl.Err(err))
			render.JSON(w, r, resp.Error("failed to decode request"))
			return
		}

		log.Info("Request body decoded", slog.Any("Request", req))

		user := &model.User{
			Email:    req.Email,
			Password: req.Password,
		}

		err = user.BeforeCreate()
		if err != nil { 
			log.Error("failed to prepare user", sl.Err(err))
			render.JSON(w, r, resp.Error("invalid user data"))
			return
		}

		log.Info("Creating user with data", slog.String("Email", user.Email), slog.String("Password", user.Password))

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
			PercentBetter: stat.PersentBetter,
		}

		render.JSON(w, r, response)
	}
}
