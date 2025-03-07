package storage

import "bench_press_calculator/internal/model"

type UserRepository interface {
	Create(u *model.User) error 
	GetAverage() (float32,error)
	UpdateWeight(userID int, weight float32) error
}
