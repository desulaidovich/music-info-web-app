package main

import (
	"os"
	"os/signal"
	"syscall"

	"app/app"
	"app/config"
	"app/internal/logger"
)

// @title Music Info
// @version 0.0.1
// description  Music Info

// @host localhost:8080
// @BasePath /

func main() {
	cfg, err := config.NewConfigFrom()
	if err != nil {
		panic(err)
	}

	app := app.NewApp(
		app.WithConfig(cfg),
		app.WithMigrate(true),
	)

	log := logger.NewLogger()

	go func() {
		log.Info(cfg.App.Name + " listening...")

		if err := app.Run(log); err != nil {
			panic(err)
		}

		log.Info("Stopped serving connections.")
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	log.Info("Shutdown complete.")
}
