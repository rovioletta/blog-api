package article

import (
	"log/slog"
	"rovioletta/blog-api/internal/database"
)

type ArticleServer struct {
	logger *slog.Logger
	db     *database.DB
}

func InitArticleServer(logger *slog.Logger, db *database.DB) *ArticleServer {
	return &ArticleServer{
		logger: logger,
		db:     db,
	}
}
