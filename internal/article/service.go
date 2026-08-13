package article

import (
	"rovioletta/blog-api/internal/db"
)

type ArticleService struct {
	db *db.DB
}

func NewArticleService(db *db.DB) *ArticleService {
	return &ArticleService{
		db: db,
	}
}
