package article

import (
	"context"
	"log/slog"

	"rovioletta/blog-api/internal/model"
)

type ArticleServiceInterface interface {
	CreateArticle(ctx context.Context, article *model.Article) error
}

type ArticleAPI struct {
	logger *slog.Logger
	server ArticleServiceInterface
}

func NewArticleAPI(logger *slog.Logger, server ArticleServiceInterface) *ArticleAPI {
	return &ArticleAPI{
		logger: logger,
		server: server,
	}
}
