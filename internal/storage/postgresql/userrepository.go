package postgresql

import (
	"bench_press_calculator/internal/lib/calc"
	"bench_press_calculator/internal/model"
	"database/sql"
	"fmt"
)

type UserRepository struct {
	store *Storage
}

func (r *UserRepository) CreateUser(u *model.User) error {
	const op = "storage.postgresql.userrepository.Create"

	err := r.store.db.QueryRow("INSERT INTO users (email,encrypted_password) VALUES ($1, $2) RETURNING id",

		u.Email,

		u.EncryptedPassword,
	).Scan(&u.ID)

	if err != nil {
		return fmt.Errorf("%s:%w", op, err)
	}
	return nil
}

func (r *UserRepository) Calculate(user *model.User, weight float32, quantity float32) (*model.Stat, error) {
	const op = "storage.postgresql.userrepository.Calculate"
	var avg sql.NullFloat64
	err := r.store.db.QueryRow("SELECT AVG(weight) FROM users").Scan(&avg)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	averageWeight := float32(1)
	if avg.Valid {
		averageWeight = float32(avg.Float64)
	}
	stat := calc.CountCalc(weight, quantity, averageWeight)
	user.Weight = stat.MaxPress
	_, err = r.store.db.Exec("UPDATE users SET weight = $1 WHERE id = $2", int32(user.Weight), user.ID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return stat, nil
}
