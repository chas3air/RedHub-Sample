package interfaces

import "net/http"

type Auth interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	Permissions(w http.ResponseWriter, r *http.Request)
}
