package article

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (a *ArticleAPI) DeleteArticleByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid query path", http.StatusBadRequest)
		return
	}

	err = a.server.DeleteArticleByID(context.Background(), id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
