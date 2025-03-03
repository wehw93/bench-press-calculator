package storage

import "bench_press_calculator/internal/model"

type UserRepository interface {
	Create(*model.User) error
	GetAverage(user *model.User) (float32, error)
	UpdateWeight(userID int, weight float32) error
}
