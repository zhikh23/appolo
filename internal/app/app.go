package app

import (
	"appolo-register/internal/app/grpcapp"
	"appolo-register/internal/app/inits"
	"appolo-register/internal/config"
	"appolo-register/internal/services/materials"
	"appolo-register/internal/storage/pgstorage"
	"log/slog"
)

type App struct {
	GrpcServer *grpcapp.App
}

func New(
	log *slog.Logger,
	cfg *config.Config,
) (*App, error) {
	db, err := inits.InitPostgres(&cfg.Postgres)
	if err != nil {
		return nil, err
	}
	storage := pgstorage.New(db)
	log.Info("Postgres storage connected")

	grpcApp := grpcapp.New(log, *materials.New(log, storage), cfg.Server.Port)

	return &App{
		GrpcServer: grpcApp,
	}, nil
}