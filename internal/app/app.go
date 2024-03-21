package app

import (
	grpcApp "PostService/internal/app/grpcApp"
	"PostService/internal/services/PostServer"
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

	server := PostServer.New(log, storage)

	grpcApplication := grpcApp.New(log, server, grpcPort)

	return &App{
		GRPCSrv: grpcApplication,
	}
}
