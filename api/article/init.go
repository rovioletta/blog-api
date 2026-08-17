package article

import (
	"context"
	"log/slog"
	"regexp"

	"github.com/go-playground/validator/v10"
	"rovioletta/blog-api/internal/model"
)

type ArticleServiceInterface interface {
	CreateArticle(ctx context.Context, article *model.Article) error
	GetArticlesByFilter(ctx context.Context, filter *model.ArticleFilter) ([]model.Article, error)
}

type ArticleAPI struct {
	logger    *slog.Logger
	validator *validator.Validate
	server    ArticleServiceInterface
}

func ValidateTagChars(fl validator.FieldLevel) bool {
	var tagCharRegex = regexp.MustCompile(`^[a-zA-Z0-9]+(-[a-zA-Z0-9]+)*$`)
	return tagCharRegex.MatchString(fl.Field().String())
}

func NewArticleAPI(logger *slog.Logger, server ArticleServiceInterface) *ArticleAPI {
	var v = validator.New(validator.WithRequiredStructEnabled())
	_ = v.RegisterValidation("tag_chars", ValidateTagChars)

	return &ArticleAPI{
		validator: v,
		logger:    logger,
		server:    server,
	}
}
