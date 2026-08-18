package article

import (
	"context"
	"fmt"
)

func (s *ArticleService) DeleteArticleByID(ctx context.Context, id uint64) error {
	err := s.db.DeleteArticleByID(ctx, id)
	if err != nil {
		return fmt.Errorf("s.db.DeleteArticle: %v", err)
	}

	return nil
}
