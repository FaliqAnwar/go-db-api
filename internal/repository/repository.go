package repository

import (
	"fmt"
	"go-db-api/internal/model"
	"go-db-api/internal/port"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(cfg model.DB) port.Repository {
	db := new(cfg)
	return repository{
		db: db,
	}
}

func new(cfg model.DB) *gorm.DB {
	connStr := fmt.Sprintf(`host=%s port=%s dbname=%s user=%s password=%s sslmode=%s`,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.Username,
		cfg.Password,
		cfg.SslMode,
	)

	connCfg, err := pgx.ParseConfig(connStr)
	if err != nil {
		log.Println(err)
	}

	sqlDB := stdlib.OpenDB(*connCfg)
	if err = sqlDB.Ping(); err != nil {
		log.Println(err)
	}

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}

	return db
}
