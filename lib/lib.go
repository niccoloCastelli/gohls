package lib

import (
	"github.com/shimberger/gohls/internal/api"
	"github.com/shimberger/gohls/internal/config"
)

type Config config.Config

func RunServer(cfg Config) {
	api.Setup(&config.Config{Folders: cfg.Folders})
}
