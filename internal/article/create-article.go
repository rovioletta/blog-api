package article

import (
	"context"
	"fmt"

	"rovioletta/blog-api/internal/db"
	"rovioletta/blog-api/internal/model"
)

func (s *ArticleService) CreateArticle(ctx context.Context, article *model.Article) error {
	err := s.db.CreateArticle(ctx, &db.CreateArticleParams{
		Title:   *article.Title,
		Content: *article.Content,
		Tags:    article.Tags,
	})

	if err != nil {
		return fmt.Errorf("s.db.CreateArticle: %v", err)
	}

	return nil
}
