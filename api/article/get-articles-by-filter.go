package article

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"rovioletta/blog-api/internal/model"
	"rovioletta/blog-api/pkg/validator"
)

type Filter struct {
	SearchTitle string    `json:"search_title" validate:"min=3,max=100"`
	SearchTags  []string  `json:"search_tags" validate:"max=5,unique,dive,alphanum,min=3,max=20"`
	CreatedFrom time.Time `json:"created_from"`
	CreatedTo   time.Time `json:"created_to"`
	UpdatedFrom time.Time `json:"updated_from"`
	UpdatedTo   time.Time `json:"updated_to"`
}

type Pagination struct {
	Limit  uint64 `json:"limit"`
	Offset uint64 `json:"offset"`
}

type Sort struct {
	Field string `json:"field" validate:"required,oneof=title created_at updated_at"`
	Order string `json:"order" validate:"required,oneof=asc desc"`
}

type GetArticlesByFilterRequestBody struct {
	Filter     *Filter     `json:"filter"`
	Pagination *Pagination `json:"pagination"`
	Sort       []Sort      `json:"sort"`
}

func (a *ArticleAPI) GetArticlesByFilter(w http.ResponseWriter, r *http.Request) {
	body, err := validator.DecodeAndValidateRequest[GetArticlesByFilterRequestBody](r, a.validator)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: get articles
	articles, err := a.server.GetArticlesByFilter(context.Background(), &model.ArticleFilter{
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := json.Marshal(articles)
	if err != nil {
		http.Error(w, "failed to encode json response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		a.logger.Error("failed to write to response", err)
	}
}
