package postgresql

import (
	"bench_press_calculator/internal/model"
	"bench_press_calculator/internal/storage"
	"database/sql"
	"fmt"
)

type UserRepository struct {
	store *Storage
}

func (r *UserRepository) Create(u *model.User) error {
	const op = "storage.postgresql.userrepository.Create"

	err := r.store.db.QueryRow("INSERT INTO users (email,encrypted_password,weight) VALUES ($1, $2, $3) RETURNING id",

		u.Email,

		u.EncryptedPassword,

		u.Weight).Scan(&u.ID)

	if err != nil {
		return fmt.Errorf("%s:%w", op, err)
	}
	return nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	const op = "storage.postgresql.userrepository.Create"

	u := &model.User{}

	if err := r.store.db.QueryRow("SELECT id,email,encrypted_password,weight FROM users WHERE email = ?",
		email).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
		&u.Weight); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("%s :%w", op, storage.ErrRecordNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return u, nil
}

func (u *UserRepository) Calculate(*model.User) (model.Stat, error){
	
}
