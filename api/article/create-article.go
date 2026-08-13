package article

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"rovioletta/blog-api/internal/model"
)

type CreateArticleRequestBody struct {
	Title   string   `json:"title" validate:"required,min=3,max=100"`
	Content string   `json:"content" validate:"required"`
	Tags    []string `json:"tags" validate:"max=5,unique,dive,alphanum,min=3,max=20"`
}

func (a *ArticleAPI) CreateArticle(w http.ResponseWriter, r *http.Request) {
	body, err := a.validateCreateArticleRequestBody(r)
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

func (a *ArticleAPI) validateCreateArticleRequestBody(r *http.Request) (body *CreateArticleRequestBody, err error) {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&body); err != nil {
		return nil, fmt.Errorf("failed to parse body: %w", err)
	}

	if err := a.validator.Struct(body); err != nil {
		return nil, fmt.Errorf("failed to validate body: %w", err)
	}

	return body, nil
}
