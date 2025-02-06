package postgresql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func New(connString string) (*Storage, error) {
	const op = "storage.postgreql.New"

	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("%s :%w", op, err)
	}
	if err:=db.Ping(); err!=nil{
		return nil,fmt.Errorf("%s :%w", op, err)
	}
	return &Storage{db: db},nil
}

func(s*Storage) Close(){
	s.db.Close()
} 

func 
