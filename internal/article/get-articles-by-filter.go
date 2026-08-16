package article

import (
	"context"
	"fmt"

	"rovioletta/blog-api/internal/db"
	"rovioletta/blog-api/internal/model"
)

func (s *ArticleService) GetArticlesByFilter(ctx context.Context, filter *model.ArticleFilter) ([]model.Article, error) {
	raw, err := s.db.GetArticlesByFilter(ctx, &db.GetArticlesByFilterParams{
		FilterSearchTitle: filter.Filter.SearchTitle,
		FilterSearchTags:  filter.Filter.SearchTags,
		FilterCreatedFrom: filter.Filter.CreatedFrom,
		FilterCreatedTo:   filter.Filter.CreatedTo,
		FilterUpdatedFrom: filter.Filter.UpdatedFrom,
		FilterUpdatedTo:   filter.Filter.UpdatedTo,
		OrderByField:      filter.Sort.Field,
		OrderType:         filter.Sort.Order,
		OffsetPage:        filter.Pagination.Offset,
		LimitPage:         filter.Pagination.Limit,
	})

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
