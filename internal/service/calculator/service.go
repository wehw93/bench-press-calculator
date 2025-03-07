package calculator

import (
	"bench_press_calculator/internal/model"
	"bench_press_calculator/internal/storage"
	"fmt"
	"log/slog"

	"github.com/labstack/gommon/log"
)

type Service struct {
	store storage.Store
}

func NewService(store storage.Store) *Service {
	return &Service{
		store: store,
	}
}

func (s *Service) Calculate(user *model.User, weight, quantity float32) (*model.Stat, error) {
	log.Info("calculate", slog.Any("user", user), slog.Float64("weight", float64(weight)), slog.Float64("quantity", float64(quantity)))

	const op = "service.calculator.Calculate"

	if err := s.store.User().Create(user); err != nil {
		return nil, fmt.Errorf("%s:%w", op, err)
	}
	avgWeight, err := s.store.User().GetAverage()
	if err != nil {
		return nil, fmt.Errorf("%s:%w", op, err)
	}
	log.Info("avg", slog.Any("AVG", avgWeight))

	stats := s.calculateMaxPress(weight, quantity, avgWeight)

	if err := s.store.User().UpdateWeight(user.ID, stats.MaxPress); err != nil {
		return nil, fmt.Errorf("%s:%w", op, err)
	}

	return stats, nil
}

func (s *Service) calculateMaxPress(weight float32, quantity float32, averageWeight float32) *model.Stat {
	MaxPress1 := (weight*quantity)/30 + weight
	MaxPress2 := weight * (1 + 0.0333*quantity)
	MaxPress3 := weight / (1.0278 - 0.0278*quantity)
	MaxPress4 := weight * 100 / (101.3 - 2.67123*quantity)
	MaxPress5 := weight * (1 + 0.025*quantity)
	MaxPress := MaxPress1 + MaxPress2 + MaxPress3 + MaxPress4 + MaxPress5
	MaxPress = MaxPress / 5
	PersentBetter := float32(100)
	if averageWeight != 0 {
		PersentBetter = ((MaxPress - averageWeight) / averageWeight) * 100
	}
	return &model.Stat{
		MaxPress:      MaxPress,
		PersentBetter: PersentBetter,
	}
}
