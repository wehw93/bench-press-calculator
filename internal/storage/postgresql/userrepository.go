package postgresql

import (
	"bench_press_calculator/internal/model"
	"database/sql"
	"fmt"
)

type UserRepository struct {
	store *Storage
}

func (r *UserRepository) Create(u *model.User) error {
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

func (r *UserRepository) GetAverage() (float32, error) {
	const op = "storage.postgresql/userrepository.GetAverage"

	var avg sql.NullFloat64
	err := r.store.db.QueryRow("SELECT AVG(weight) FROM users").Scan(&avg)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	if !avg.Valid {
		return 0, nil
	}

	return float32(avg.Float64), nil
}

func (r *UserRepository) UpdateWeight(userID int, weight float32) error {
	const op = "storage.postgresql.userrepository.UpdateWeight"

	tx, err := r.store.db.Begin()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if _, err := tx.Exec(`UPDATE users SET weight = $1 WHERE id = $2`, weight, userID); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return tx.Commit()
}
