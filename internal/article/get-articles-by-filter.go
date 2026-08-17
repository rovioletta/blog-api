package article

import (
	"context"
	"fmt"

	"rovioletta/blog-api/internal/db"
	"rovioletta/blog-api/internal/model"
)

func (s *ArticleService) GetArticlesByFilter(ctx context.Context, filter *model.ArticleFilter) ([]model.Article, error) {
	if filter == nil {
		return nil, fmt.Errorf("no data")
	}

	params := &db.GetArticlesByFilterParams{}

	if filter.Filter != nil {
		params.FilterSearchTitle = filter.Filter.SearchTitle
		params.FilterSearchTags = filter.Filter.SearchTags
		params.FilterCreatedFrom = filter.Filter.CreatedFrom
		params.FilterCreatedTo = filter.Filter.CreatedTo
		params.FilterUpdatedFrom = filter.Filter.UpdatedFrom
		params.FilterUpdatedTo = filter.Filter.UpdatedTo
	}

	if filter.Sort != nil {
		params.OrderByField = filter.Sort.Field
		params.OrderType = filter.Sort.Order
	}

	if filter.Pagination != nil {
		params.Offset = filter.Pagination.Offset
		params.Limit = filter.Pagination.Limit
	}

	raw, err := s.db.GetArticlesByFilter(ctx, params)

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
