package main

import (
	"os"

	ht "github.com/samarec1812/segmentation-service/internal/app/ports/http"
	"github.com/samarec1812/segmentation-service/internal/config"
	"github.com/samarec1812/segmentation-service/internal/pkg/logger"
	"github.com/samarec1812/segmentation-service/internal/pkg/postgres"
)

func main() {

	//ctx := context.Background()
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Env)
	log.Info("starting application")
	log.Debug("debug message")

	_, err := postgres.Connect(cfg.DB_URL)
	if err != nil {
		log.Error("error connect database", err)
		os.Exit(1)
	}

	log.Info("database connect successful")
	a := ht.NewApp()
	srv := ht.NewHTTPServer(cfg.Address, log, a)

	if err = srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
		os.Exit(1)
	}
}
