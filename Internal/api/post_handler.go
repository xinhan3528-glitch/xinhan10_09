package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"xinhan10_09/Internal/store"

	"github.com/go-chi/chi/v5"
)

type PostHandler struct {
	postStore store.PostStore
}

func NewPostHandler(ps store.PostStore) *PostHandler {
	return &PostHandler{postStore: ps}
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

	var post store.Post

	err := json.NewDecoder(r.Body).Decode(&post) //change json to golang
	if err != nil {
		http.Error(w, "Invaild request data", http.StatusBadRequest)
		return
	} // to here..

	//now add to database..

	createdPost, err := ph.postStore.CreatePost(&post) //&post = var post store.Post; and to createPost database process
	if err != nil {
		fmt.Println(err)
		http.Error(w, "can't create Post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)      //assign w.Header
	json.NewEncoder(w).Encode(createdPost) //golang to json

}
