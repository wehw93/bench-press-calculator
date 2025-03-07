package service

import "bench_press_calculator/internal/model"

type CalculatorService interface {
	Calculate(user *model.User, weight, quantity float32) (*model.Stat, error)
}
