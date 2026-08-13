package article

import (
	"net/http"
)

func (s *ArticleAPI) CreateArticle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("success"))

	// Insert data
	// TODO
}
