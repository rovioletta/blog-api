package article

import (
	"context"
	"errors"
	"fmt"

	"rovioletta/blog-api/internal/model"

	"github.com/jackc/pgx/v5"
)

func (s *ArticleService) GetArticlesByID(ctx context.Context, id uint64) (*model.Article, error) {
	raw, err := s.db.GetArticleByID(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, model.ErrNoData
	}

	if err != nil {
		return nil, fmt.Errorf("s.db.GetArticleByID: %v", err)
	}

	return &model.Article{
		ID:      raw.ID,
		Title:   raw.Title,
		Content: raw.Content,
		Tags:    raw.Tags,
	}, nil
}
