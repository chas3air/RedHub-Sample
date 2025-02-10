package users

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"usersManageService/internal/config"
	"usersManageService/internal/domain/models"
	"usersManageService/lib/logger/sl"

	"github.com/gorilla/mux"
)

type UserManageController struct {
	log    *slog.Logger
	client *http.Client
}

func New(logger *slog.Logger, client *http.Client) *UserManageController {
	return &UserManageController{
		log:    logger,
		client: client,
	}
}

func handleResponseErrors(u *UserManageController, resp *http.Response, w http.ResponseWriter) bool {
	if resp.StatusCode == http.StatusServiceUnavailable {
		u.log.Error("service unavailable", slog.Int("code", resp.StatusCode))
		http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
		return true
	}

	if resp.StatusCode == http.StatusBadRequest {
		u.log.Error("resource not found", slog.Int("code", resp.StatusCode))
		http.Error(w, "Not Found", http.StatusNotFound)
		return true
	}

	if resp.StatusCode >= 400 {
		u.log.Error("status code is unavailable", slog.Int("code", resp.StatusCode))
		http.Error(w, resp.Status, resp.StatusCode)
		return true
	}

	return false
}

func (u *UserManageController) Get(w http.ResponseWriter, r *http.Request) {
	u.log.Info("/users (GET) running...")

	resp, err := u.client.Get(config.UsersAccessService_url)
	if err != nil {
		u.log.Error("/users (GET), error", sl.Err(err))
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	if handleResponseErrors(u, resp, w) {
		return
	}

	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		u.log.Error("cannot read response body", sl.Err(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(bs); err != nil {
		u.log.Error("error encoding response", sl.Err(err))
		http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	u.log.Info("/users (GET) done")
}

func (u *UserManageController) GetById(w http.ResponseWriter, r *http.Request) {
	u.log.Info("/users/{id}/ (GET) running...")
	id := mux.Vars(r)["id"]
	if id == "" {
		http.Error(w, "Id is required", http.StatusBadRequest)
		return
	}

	resp, err := u.client.Get("URL" + id + "/")
	if err != nil {
		u.log.Error("/users/{id}/ (GET), error", sl.Err(err))
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	if handleResponseErrors(u, resp, w) {
		return
	}

	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		u.log.Error("cannot read response body", sl.Err(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(bs); err != nil {
		u.log.Error("error encoding response", sl.Err(err))
		http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	u.log.Info("/users/{id}/ (GET) done")
}

// GetByNick implements interfaces.UserManager.
func (u *UserManageController) GetByNick(w http.ResponseWriter, r *http.Request) {
	u.log.Info("/users/{nick} (GET) running...")
	nick := mux.Vars(r)["nick"]
	if nick == "" {
		http.Error(w, "Nick is required", http.StatusBadRequest)
		return
	}

	resp, err := u.client.Get("URL")
	if err != nil {
		u.log.Error("/users (GET), error", sl.Err(err))
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	if handleResponseErrors(u, resp, w) {
		return
	}

	var users []models.User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		u.log.Error("cannot marshal response body to objects", sl.Err(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var neededUser models.User
	for _, user := range users {
		if user.Nick == nick {
			neededUser = user
			break
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(neededUser); err != nil {
		u.log.Error("error encoding response", sl.Err(err))
		http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	u.log.Info("/users/{nick} (GET) done")
}

func (u *UserManageController) Insert(w http.ResponseWriter, r *http.Request) {
	u.log.Info("/users (POST) running...")

	resp, err := u.client.Post("URL", "application/json", r.Body)
	if err != nil {
		u.log.Error("/users (POST), error", sl.Err(err))
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	if handleResponseErrors(u, resp, w) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	u.log.Info("/users (POST) done")
}

func (u *UserManageController) Update(w http.ResponseWriter, r *http.Request) {
	u.log.Info("/users/{id} (PUT) running...")
	id := mux.Vars(r)["id"]
	if id == "" {
		http.Error(w, "Id is required", http.StatusBadRequest)
		return
	}

	req, err := http.NewRequest(http.MethodPut, "URL/"+id, r.Body)
	if err != nil {
		u.log.Error("cannot create request", sl.Err(err))
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	resp, err := u.client.Do(req)
	if err != nil {
		u.log.Error("request execution error", sl.Err(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if handleResponseErrors(u, resp, w) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	u.log.Info("/users/{id} (PUT) done")
}

func (u *UserManageController) Delete(w http.ResponseWriter, r *http.Request) {
	u.log.Info("/users/{id} (DELETE) running...")
	id := mux.Vars(r)["id"]
	if id == "" {
		http.Error(w, "Id is required", http.StatusBadRequest)
		return
	}

	req, err := http.NewRequest(http.MethodDelete, "URL/"+id, nil)
	if err != nil {
		u.log.Error("cannot create request", sl.Err(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := u.client.Do(req)
	if err != nil {
		u.log.Error("request execution error", sl.Err(err))
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	if handleResponseErrors(u, resp, w) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp.Body); err != nil {
		u.log.Error("error encoding response", sl.Err(err))
		http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
		return
	}
	u.log.Info("/users/{id} (DELETE) done")
}
