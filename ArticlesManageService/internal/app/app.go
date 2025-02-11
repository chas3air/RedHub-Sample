package app

import (
	"articlesManageService/internal/config"
	"articlesManageService/internal/controllers/articles"
	"articlesManageService/internal/domein/interfaces"
	"articlesManageService/lib/logger/sl"
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

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
	var articlescontroller interfaces.ArticlesManager = articles.New(a.log, &http.Client{Timeout: a.cfg.ExpirationTime})

	r := mux.NewRouter()

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	r.HandleFunc("/articles", articlescontroller.Get).Methods(http.MethodGet)
	r.HandleFunc("/articles/{id}", articlescontroller.GetById).Methods(http.MethodGet)
	r.HandleFunc("/articles", articlescontroller.Insert).Methods(http.MethodPost)
	r.HandleFunc("/articles/{id}", articlescontroller.Update).Methods(http.MethodPut)
	r.HandleFunc("/articles/{id}", articlescontroller.Delete).Methods(http.MethodDelete)
	r.HandleFunc("/articles/{postId}/coments", articlescontroller.InteractionWithCommentsOfPost).Methods(http.MethodGet, http.MethodPost)

	a.srv = &http.Server{
		Addr:    fmt.Sprintf(":%d", a.cfg.Port),
		Handler: r,
	}

	a.wg.Add(1)
	go func() {
		defer a.wg.Done()
		if err := a.srv.ListenAndServe(); err != nil {
			a.log.Error("error of starting server", sl.Err(err))
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	return a.Stop()
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
