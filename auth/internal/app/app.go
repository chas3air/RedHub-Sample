package app

import (
	"auth/internal/config"
	"auth/internal/controllers/auth"
	"auth/internal/domein/interfaces"
	"auth/lib/logger/sl"
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"

	"github.com/gin-gonic/gin"
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
	var authcontroller interfaces.Auth = auth.New(a.log, &http.Client{Timeout: a.cfg.ExpirationTime})

	router := gin.Default()
	router.GET("/auth/login", authcontroller.Login)
	router.POST("/auth/register", authcontroller.Register)
	router.GET("/auth/permissions", authcontroller.Permissions)

	a.srv = &http.Server{
		Addr:    fmt.Sprintf(":%d", a.cfg.Port),
		Handler: router,
	}

	a.wg.Add(1)
	go func() {
		defer a.wg.Done()
		if err := router.Run(":" + strconv.Itoa(a.cfg.Port)); err != nil {
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
