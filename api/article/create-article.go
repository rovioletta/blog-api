package article

import (
	"context"
	"net/http"

	"rovioletta/blog-api/internal/model"
	"rovioletta/blog-api/pkg/validator"
)

type CreateArticleRequestBody struct {
	Title   string   `json:"title" validate:"required,min=3,max=100"`
	Content string   `json:"content" validate:"required"`
	Tags    []string `json:"tags" validate:"max=5,unique,dive,alphanum,min=3,max=20"`
}

func (a *ArticleAPI) CreateArticle(w http.ResponseWriter, r *http.Request) {
	body, err := validator.DecodeAndValidateRequest[CreateArticleRequestBody](r, a.validator)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = a.server.CreateArticle(context.Background(), &model.Article{
		Title:   body.Title,
		Content: body.Content,
		Tags:    body.Tags,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
