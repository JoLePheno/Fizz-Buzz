package postgres

import (
	"fmt"
	"log"

	"github.com/caarlos0/env"
	"github.com/go-pg/pg"
)

var cfg Config

type Config struct {
	Port     int    `env:"POSTGRES_PORT" envDefault:"5432"`
	Host     string `env:"POSTGRES_HOST" envDefault:"postgres"`
	Password string `env:"POSTGRES_PASSWORD" envDefault:"password"`
	User     string `env:"POSTGRES_USER" envDefault:"postgres"`
	Database string `env:"POSTGRES_DB" envDefault:"postgres"`
}

func NewPostgresStore() *ParametersStore {
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatalf("Failed to parse postgres environment: %v", err)
	}

	db := pg.Connect(&pg.Options{
		User:     cfg.User,
		Database: cfg.Database,
		Password: cfg.Password,
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
	})
	return &ParametersStore{db: db}
}

func NewPostgres() *pg.DB {
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatalf("Failed to parse postgres environment: %v", err)
	}

	db := pg.Connect(&pg.Options{
		User:     cfg.User,
		Database: cfg.Database,
		Password: cfg.Password,
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
	})

	return db
}

func (s *ParametersStore) ClearStore() {
	s.db.Exec(`TRUNCATE TABLE parameters`)
}