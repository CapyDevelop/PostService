package app

import (
	grpcApp "PostService/internal/app/grpcApp"
	storage2 "PostService/internal/storage"
	"PostService/internal/storage/postgresConn"
	"log/slog"
)

type App struct {
	GRPCSrv *grpcApp.App
}

func New(log *slog.Logger, grpcPort int,
	Username, Password,
	Database,
	Host string,
	Port int) *App {

	dbConn, err := postgresConn.New(Username, Password, Database, Host, Port)
	if err != nil {
		panic(err)
	}

	storage := storage2.New(log, dbConn)

	grpcApplication := grpcApp.New(log, storage, grpcPort)

	return &App{
		GRPCSrv: grpcApplication,
	}
}
