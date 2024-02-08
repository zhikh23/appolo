package inits

import (
	"appolo-register/internal/config"
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/QuickDrone-Backend/pgconn"
	"github.com/jmoiron/sqlx"
)

const (
	dbConnectAttempts = 3
	dbConnectRetryDuration = time.Second
)

func InitPostgres(cfg *config.Postgres) (*sqlx.DB, error) {
	connConfig := pgconn.ConnConfig{
		Host:     cfg.Host,
		Port:     cfg.Port,
		User:     cfg.User,
		Password: cfg.Password,
		DbName:   cfg.DbName,
		SslMode:  cfg.SslMode,
	}
	db, err := pgconn.ConnectWithTries(
		context.Background(),
		connConfig.Url(),
		nil,
		dbConnectAttempts,
		dbConnectRetryDuration,
	)
	if err != nil {
		return nil, err
	}
	slog.Debug(
		"Postgres successfully started",
		slog.String("host:port", fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)),
	)
	return db, nil
}