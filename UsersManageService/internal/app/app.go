package app

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"usersManageService/internal/config"
	"usersManageService/internal/controllers/users"
	"usersManageService/internal/domain/interfaces"

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
	var userscontroller interfaces.UserManager = users.New(a.log, &http.Client{Timeout: a.cfg.ExpirationTime})

	r := mux.NewRouter()
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	r.HandleFunc("/users", userscontroller.Get).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}/", userscontroller.GetById).Methods(http.MethodGet)
	r.HandleFunc("/users/{nick}", userscontroller.GetByNick).Methods(http.MethodGet)
	r.HandleFunc("/users", userscontroller.Insert).Methods(http.MethodPost)
	r.HandleFunc("/users/{id}", userscontroller.Update).Methods(http.MethodPut)
	r.HandleFunc("/users/{id}", userscontroller.Delete).Methods(http.MethodDelete)

	a.srv = &http.Server{
		Addr:    fmt.Sprintf(":%d", a.cfg.Port),
		Handler: r,
	}

	a.wg.Add(1)
	go func() {
		defer a.wg.Done()
		if err := a.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.log.Error("error of starting server", "error", err)
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
