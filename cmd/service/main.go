package main

import (
	"fmt"
	"github.com/samarec1812/segmentation-service/internal/config"
	"github.com/samarec1812/segmentation-service/internal/pkg/logger"
	"github.com/samarec1812/segmentation-service/internal/pkg/postgres"
	"os"
)

func main() {

	//ctx := context.Background()
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Env)
	log.Info("starting application")
	log.Debug("debug message")
	fmt.Println(cfg)

	_, err := postgres.Connect(cfg.DB_URL)
	if err != nil {
		log.Error("error connect database", err)
		os.Exit(1)
	}

	log.Info("database connect successful")
}
