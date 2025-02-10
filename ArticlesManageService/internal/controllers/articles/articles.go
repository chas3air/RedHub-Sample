package articles

import (
	"articlesManageService/internal/config"
	"articlesManageService/lib/logger/sl"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

type ArticlesManageService struct {
	log    *slog.Logger
	client *http.Client
}

func New(logger *slog.Logger, client *http.Client) *ArticlesManageService {
	return &ArticlesManageService{
		log:    logger,
		client: client,
	}
}

func (a *ArticlesManageService) Get(w http.ResponseWriter, r *http.Request) {
	a.log.Info("/articles (GET) running...")
	resp, err := a.client.Get(config.ArticlesAccessService_url)
	if err != nil {
		a.log.Error("/articles (GET)", sl.Err(err))
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	if handleResponseErrors(a, resp, w) {
		return
	}

	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		a.log.Error("cannot read response body", sl.Err(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(bs); err != nil {
		a.log.Error("error encoding response", sl.Err(err))
		http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	a.log.Info("/articles (GET) done")
}

func (a *ArticlesManageService) GetById(w http.ResponseWriter, r *http.Request) {
	a.log.Info("/articles/{id} (GET) running...")
	id := mux.Vars(r)["id"]
	if id == "" {
		a.log.Error("Id is required")
		http.Error(w, "Id is required", http.StatusBadRequest)
		return
	}

	resp, err := a.client.Get(config.ArticlesAccessService_url)
	if err != nil {
		a.log.Error("/articles/{id} (GET)", sl.Err(err))
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	if handleResponseErrors(a, resp, w) {
		return
	}

	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		a.log.Error("cannot read response body", sl.Err(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(bs); err != nil {
		a.log.Error("error encoding response", sl.Err(err))
		http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	a.log.Info("/articles/{id} (GET) done")
}

func (a *ArticlesManageService) Insert(w http.ResponseWriter, r *http.Request) {
	a.log.Info("/articles (POST) running...")

	resp, err := a.client.Post(config.ArticlesAccessService_url, "application/json", r.Body)
	if err != nil {
		a.log.Error("/articles (POST)", sl.Err(err))
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	if handleResponseErrors(a, resp, w) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	a.log.Info("/articles (POST) done")
}

func (a *ArticlesManageService) Update(w http.ResponseWriter, r *http.Request) {
	a.log.Info("/articles/{id} (PUT) running...")
	id := mux.Vars(r)["id"]
	if id == "" {
		a.log.Error("Id is required")
		http.Error(w, "Id is required", http.StatusBadRequest)
		return
	}

	req, err := http.NewRequest(http.MethodPut, config.ArticlesAccessService_url+"/"+id, r.Body)
	if err != nil {
		a.log.Error("cannot create request:", sl.Err(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := a.client.Do(req)
	if err != nil {
		a.log.Error("/articles (PUT)", sl.Err(err))
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	if handleResponseErrors(a, resp, w) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	a.log.Info("/articles/{id} (PUT) done")
}

func (a *ArticlesManageService) Delete(w http.ResponseWriter, r *http.Request) {
	a.log.Info("/articles/{id} (DELETE) running...")
	id := mux.Vars(r)["id"]
	if id == "" {
		a.log.Error("Id is required")
		http.Error(w, "Id is required", http.StatusBadRequest)
		return
	}

	req, err := http.NewRequest(http.MethodDelete, config.ArticlesAccessService_url+"/"+id, nil)
	if err != nil {
		a.log.Error("cannot create request", sl.Err(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := a.client.Do(req)
	if err != nil {
		a.log.Error("/articles (DELETE)", sl.Err(err))
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	if err := json.NewEncoder(w).Encode(resp.Body); err != nil {
		a.log.Error("error encoding response", sl.Err(err))
		http.Error(w, "Failed to encode response:"+err.Error(), http.StatusInternalServerError)
		return
	}
	a.log.Info("/articles/{id} (DELETE) done")
}

func handleResponseErrors(a *ArticlesManageService, resp *http.Response, w http.ResponseWriter) bool {
	if resp.StatusCode == http.StatusServiceUnavailable {
		a.log.Error("service unavailable", slog.Int("code", resp.StatusCode))
		http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
		return true
	}

	if resp.StatusCode == http.StatusBadRequest {
		a.log.Error("resource not found", slog.Int("code", resp.StatusCode))
		http.Error(w, "Not Found", http.StatusNotFound)
		return true
	}

	if resp.StatusCode >= 400 {
		a.log.Error("status code is unavailable", slog.Int("code", resp.StatusCode))
		http.Error(w, resp.Status, resp.StatusCode)
		return true
	}

	return false
}
