package app

import (
	"api-gateway/internal/config"
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

type App struct {
	log *slog.Logger
	cfg *config.Config
	srv *http.Server
	wg  sync.WaitGroup
}

func New(log *slog.Logger, cfg *config.Config) *App {
	return &App{
		log: log,
		cfg: cfg,
	}
}

func (a *App) StartServer() error {
	// TODO: написать обработчик usercontroller
	// TODO: написать обработчик articlecontroller
	// TODO: написать обработчик commentcontroller
	// TODO: написать обработчик authcontroller

	r := mux.NewRouter()

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.HandleFunc("/users", nil).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}", nil).Methods(http.MethodGet)
	r.HandleFunc("/users/email/{email}", nil).Methods(http.MethodGet)
	r.HandleFunc("/users", nil).Methods(http.MethodPost)
	r.HandleFunc("/users/{id}", nil).Methods(http.MethodPut)
	r.HandleFunc("/users/{id}", nil).Methods(http.MethodDelete)

	r.HandleFunc("/articles", nil).Methods(http.MethodGet)
	r.HandleFunc("/articles/{id}", nil).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}/articles", nil).Methods(http.MethodGet)
	r.HandleFunc("/articles", nil).Methods(http.MethodPost)
	r.HandleFunc("/articles/{id}", nil).Methods(http.MethodPut)
	r.HandleFunc("/articles/{id}", nil).Methods(http.MethodDelete)

	r.HandleFunc("/articles/{aid}/comments", nil).Methods(http.MethodGet)
	r.HandleFunc("/articles/{aid}/comments", nil).Methods(http.MethodPost)
	r.HandleFunc("/comments/{cid}", nil).Methods(http.MethodPut)
	r.HandleFunc("/comments{cid}", nil).Methods(http.MethodDelete)

	a.srv = &http.Server{
		Addr:    fmt.Sprintf(":%d", a.cfg.Port),
		Handler: r,
	}

	go func() {
		if err := a.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.log.Error("error of starting server", "error", err)
		}
	}()

	return nil
}

func (a *App) Stop() error {
	a.log.Info("Stoping server...")

	ctx, cancel := context.WithTimeout(context.Background(), a.cfg.ExpirationTime)
	defer cancel()

	if err := a.srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("error while stoping server: %v", err)
	}

	a.wg.Wait()
	a.log.Info("Server is stoped")
	return nil
}
