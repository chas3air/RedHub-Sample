package app

import (
	"log/slog"
	grpcapp "usersManageService/internal/app/grpc"
	"usersManageService/internal/storage/mock"
	usermanager "usersManageService/services/userManager"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(log *slog.Logger, port int) *App {
	storage := mock.New()
	usermanager := usermanager.New(log, storage)

	grpcapp := grpcapp.New(log, usermanager, port)
	return &App{
		GRPCServer: grpcapp,
	}
}
