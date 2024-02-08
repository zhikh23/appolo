package main

import (
	"appolo-register/internal/app"
	"appolo-register/internal/app/inits"
	"appolo-register/internal/config"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()

	logger := inits.InitLogger(cfg.Env)
	slog.SetDefault(logger)

	application, err := app.New(logger, cfg)
	if err != nil {
		panic(err)
	}

	go func() {
		application.GrpcServer.MustRun()
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.GrpcServer.Stop()
	logger.Info("Gracefully stopped")
}
