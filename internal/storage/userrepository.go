package storage

import "bench_press_calculator/internal/model"

type UserRepository interface {
	CreateUser(*model.User) error
	FindByEmail(string) (*model.User, error)
	Calculate(*model.User, float32,float32) (*model.Stat, error)
}
