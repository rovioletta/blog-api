package article

import (
	"context"
	"log/slog"

	"github.com/go-playground/validator/v10"
	"rovioletta/blog-api/internal/model"
)

type ArticleServiceInterface interface {
	CreateArticle(ctx context.Context, article *model.Article) error
}

type ArticleAPI struct {
	logger    *slog.Logger
	validator *validator.Validate
	server    ArticleServiceInterface
}

func NewArticleAPI(logger *slog.Logger, server ArticleServiceInterface) *ArticleAPI {
	var v = validator.New(validator.WithRequiredStructEnabled())

	return &ArticleAPI{
		validator: v,
		logger:    logger,
		server:    server,
	}
}
