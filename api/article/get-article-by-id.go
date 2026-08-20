package article

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"rovioletta/blog-api/internal/model"
)

func (a *ArticleAPI) GetArticleByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid url", http.StatusBadRequest)
		return
	}

	article, err := a.server.GetArticleByID(context.Background(), id)
	if errors.Is(err, model.ErrNoData) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "failed to encode json response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		a.logger.Error("failed to write to response", slog.String("error", err.Error()))
	}
}
