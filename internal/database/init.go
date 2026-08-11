package database

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
)

type DB struct {
	logger *slog.Logger
	Conn   *pgx.Conn
}

func InitDB(logger *slog.Logger) *DB {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Error("Unable to connect to database", slog.String("error", err.Error()))
		os.Exit(1)
	}

	return &DB{
		logger: logger,
		Conn:   conn,
	}
}

func (db *DB) CloseDB() {
	db.Conn.Close(context.Background())
}
