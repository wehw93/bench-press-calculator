package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env        string     `yaml:"env" env:"ENV" env-default:"prod"`
	DB         DB         `yaml:"db"`
	HTTPServer HTTPServer `yaml:"http_server"`
}

type DB struct {
	Host     string `yaml:"host" env:"PG_HOST" env-default:"calculator_db"`
	Port     string `yaml:"port" env:"PG_PORT" env-default:"5432"`
	Name     string `yaml:"name" env:"PG_DBNAME" env-default:"calculator_db"`
	User     string `yaml:"user" env:"PG_USER" env-default:"calc_user"`
	Password string `yaml:"password" env:"PG_PASSWORD" env-default:"pwd123"`
	SSLMode  string `yaml:"sslmode" env:"PG_SSLMODE" env-default:"disable"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"0.0.0.0:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func (db *DB) GetDSN() string {
	return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		db.Host, db.Port, db.Name, db.User, db.Password, db.SSLMode)
}

func MustLoad() Config {
	if err := godotenv.Load("local.env"); err != nil && !os.IsNotExist(err) {
		log.Fatalf("error loading .env file: %v", err)
	}
	var cfg Config
	if err := cleanenv.ReadConfig(os.Getenv("CONFIG_PATH"), &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}
	return cfg
}
