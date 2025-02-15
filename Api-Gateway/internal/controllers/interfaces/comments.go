package interfaces

import "net/http"

type CommentsHandler interface {
	ListCommentsHandler(w http.ResponseWriter, r *http.Request)
	GetByIdHandler(w http.ResponseWriter, r *http.Request)
	GetByLoginHandler(w http.ResponseWriter, r *http.Request)
	Insert(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}
