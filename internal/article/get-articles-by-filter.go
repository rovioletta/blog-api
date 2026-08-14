package article

import (
	"context"
	"fmt"

	"rovioletta/blog-api/internal/db"
	"rovioletta/blog-api/internal/model"
)

func (s *ArticleService) GetArticlesByFilter(ctx context.Context, filter *model.ArticleFilter) ([]model.Article, error) {
	// TODO insert filter paraments
	// TODO add orderBy and pagination
	raw, err := s.db.GetArticlesByFilter(ctx, &db.GetArticlesByFilterParams{})

	if err != nil {
		return nil, fmt.Errorf("s.db.GetArticlesByFilter: %v", err)
	}

	articles := make([]model.Article, 0, len(raw))

	for _, a := range raw {
		articles = append(articles, model.Article{
			ID:      a.ID,
			Title:   a.Title,
			Content: a.Content,
			Tags:    a.Tags,
		})
	}

	return articles, nil
}
