package postgresql

import (
	"bench_press_calculator/internal/storage"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage struct {
	db             *sql.DB
	userRepository *UserRepository
}

func New(dsn string) (*Storage, error) {
	const op = "storage.postgreql.New"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("%s :%w", op, err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("%s :%w", op, err)
	}
	return &Storage{db: db}, nil
}

func (s *Storage) User() storage.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
	}
	return s.userRepository
}

func (s *Storage) Close() {
	s.db.Close()
}
