package interfaces

import "net/http"

type AuthHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	IsAdmin(w http.ResponseWriter, r *http.Request)
}
