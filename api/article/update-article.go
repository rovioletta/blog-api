package article

import (
	"context"
	"net/http"
	"strconv"

	"rovioletta/blog-api/internal/model"
	"rovioletta/blog-api/pkg/validator"

	"github.com/go-chi/chi"
)

type UpdateArticleRequestBody struct {
	Title   *string  `json:"title" validate:"omitempty,min=3,max=100"`
	Content *string  `json:"content"`
	Tags    []string `json:"tags" validate:"omitempty,max=5,unique,dive,tag_chars,min=3,max=20"`
}

func (a *ArticleAPI) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	body, err := validator.DecodeAndValidateRequest[UpdateArticleRequestBody](r, a.validator)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid url", http.StatusBadRequest)
		return
	}

	err = a.server.UpdateArticle(context.Background(), &model.Article{
		ID:      id,
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
