package db

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
)

type DB struct {
	*Queries
	logger *slog.Logger
	conn   *pgx.Conn
}

func NewDB(logger *slog.Logger) *DB {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Error("Unable to connect to database", slog.String("error", err.Error()))
		os.Exit(1)
	}

	return &DB{
		Queries: New(conn),
		logger:  logger,
		conn:    conn,
	}
}

func (db *DB) CloseDB() {
	db.logger.Info("Closing database connection...")
	db.conn.Close(context.Background())
}
