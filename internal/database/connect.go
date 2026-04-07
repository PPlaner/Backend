package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/PPlaner/Backend/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect(cfg config.Database) (*sql.DB, error) {
	dns := fmt.Sprintf(""+
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Name,
		cfg.SSLMode,
	)

	db, err := sql.Open("pgx", dns)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)

	if err := db.Ping(); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil

}
