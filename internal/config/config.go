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
	Env string `yaml:"env" env:"ENV" env-default:"prod"`

	DB         DB         `yaml:"db"`
	HTTPServer HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type DB struct {
	Host     string `yaml:"host" env-default:"calculator_db"`
	Port     string `yaml:"port" env-default:"5432"`
	Name     string `yaml:"name" env-default:"calculator_db"`
	User     string `yaml:"user" env-default:"calc_user"`
	Password string `yaml:"password" env-default:"pwd123"`
	SSLMode  string `yaml:"sslmode" env-default:"disable"`
}

func (db *DB) GetDSN() string {
	return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", db.Host, db.Port, db.Name, db.User, db.Password, db.SSLMode)
}

func MustLoad() *Config {
	_, err := os.Stat("local.env")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("нет файла")
		}
	}

	err = godotenv.Load("local.env")
	if err != nil {
		log.Fatal("error of Load .env")
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG PATH IS NOT SET")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exits: %s", configPath)
	}

	var cfg *Config
	if err := cleanenv.ReadConfig(configPath, cfg); err != nil {
		log.Fatalf("cannot read config %s", err)
	}

	return cfg
}
