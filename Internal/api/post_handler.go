package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type PostHandler struct{}

func NewPostHandler() *PostHandler {
	return &PostHandler{}
}

func (ph *PostHandler) HandlerGetPostById(w http.ResponseWriter, r *http.Request) {
	paramsPostId := chi.URLParam(r, "id")
	if paramsPostId == "" {
		http.NotFound(w, r)
		return
	}

	postId, err := strconv.ParseInt(paramsPostId, 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "这是Post ID %d\n", postId)
}

func (ph *PostHandler) HandlerCreatePost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "新增Post已开建\n")
}
