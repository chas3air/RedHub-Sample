package app

import (
	"log/slog"
	grpcapp "usersManageService/internal/app/grpc"
	usermanager "usersManageService/services/userManager"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(log *slog.Logger, port int) *App {
	usermanager := usermanager.New(log, nil)

	grpcapp := grpcapp.New(log, usermanager, port)
	return &App{
		GRPCServer: grpcapp,
	}
}
