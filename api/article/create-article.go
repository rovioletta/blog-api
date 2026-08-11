package article

import (
	"net/http"
)

func (s *ArticleServer) CreateArticle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("success"))

	// Insert data
}
