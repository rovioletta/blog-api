package article

import (
	"context"
	"fmt"

	"rovioletta/blog-api/internal/db"
	"rovioletta/blog-api/internal/model"
)

func (s *ArticleService) UpdateArticle(ctx context.Context, article *model.Article) error {
	err := s.db.UpdateArticle(ctx, &db.UpdateArticleParams{
		NewTitle:   article.Title,
		NewContent: article.Content,
		NewTags:    article.Tags,
		ID:         article.ID,
	})

	if err != nil {
		return fmt.Errorf("s.db.UpdateArticle: %v", err)
	}

	return nil
}
