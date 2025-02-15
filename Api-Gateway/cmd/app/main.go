package main

import (
	"api-gateway/internal/app"
	"api-gateway/internal/config"
	"api-gateway/internal/logger"
	"api-gateway/internal/logger/sl"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Env)

	application := app.New(log, cfg)

	log.Info("server is available:", slog.Int("post", cfg.Port))

	go func() {
		if err := application.StartServer(); err != nil {
			log.Error("server error", sl.Err(err))
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	log.Info("stoping application")
	if err := application.Stop(); err != nil {
		log.Error("error stoping application", sl.Err(err))
	}
	log.Info("application stoped")
}
