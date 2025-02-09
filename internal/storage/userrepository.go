package storage

import "bench_press_calculator/internal/model"

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
	Calculate(*model.User) (model.Stat, error)
}
