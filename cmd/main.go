package main

import (
	logger "github.com/dm1trypon/easy-logger"
	"github.com/dm1trypon/engine-srv/config"
	"github.com/dm1trypon/engine-srv/manager"
)

// LC - logging's category
const LC = "MAIN"

func main() {
	cfg := logger.Cfg{
		AppName: "ENGINE_SRV",
		LogPath: "",
		Level:   0,
	}

	logger.SetConfig(cfg)

	logger.Info(LC, "STARTING SERVICE")

	config.ConfigInst = new(config.Config).Create()
	new(manager.Manager).Create()
}
