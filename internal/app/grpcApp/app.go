package grpcApp

import (
	"PostService/internal/grpc/handlers"
	"fmt"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(log *slog.Logger,
	postServ handlers.PostServ,
	port int) *App {
	gRPCServer := grpc.NewServer()
	handlers.Register(gRPCServer, postServ)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "grpc.Run"

	log := a.log.With(
		slog.String("op", op),
		slog.Int("port", a.port))

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("gRPC server is running", slog.String("addr", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "grpcApp.Stop"

	a.log.With(slog.String("op", op)).
		Info("stopping gRPC server")

	a.gRPCServer.GracefulStop()

}
