package main

import (
	"github.com/samarec1812/segmentation-service/internal/app"
	"github.com/samarec1812/segmentation-service/internal/config"
)

func main() {

	//ctx := context.Background()
	cfg := config.MustLoad()

	app.Run(cfg)
}
