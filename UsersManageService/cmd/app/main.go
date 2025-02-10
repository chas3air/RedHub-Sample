package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"usersManageService/internal/app"
	"usersManageService/internal/config"
	"usersManageService/lib/logger"
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

	log.Info("stopping application")
	if err := application.Stop(); err != nil {
		log.Error("error stopping application", slog.Any("error", err))
	}
	log.Info("application stopped")
}
