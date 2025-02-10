package main

import (
	"articlesManageService/internal/app"
	"articlesManageService/internal/config"
	"articlesManageService/lib/logger"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Env)

	log.Info("starting application", slog.Any("config:", cfg))

	application := app.New(log, cfg)

	go func() {
		if err := application.StartServer(); err != nil {
			log.Error("server error", slog.Any("error", err))
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	log.Info("stoping application")
	if err := application.Stop(); err != nil {
		log.Error("error stoping application", slog.Any("error", err))
	}
	log.Info("application stopped")
}
