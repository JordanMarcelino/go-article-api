package main

import (
	"github.com/jordanmarcelino/go-article-api/internal/config"
	"github.com/jordanmarcelino/go-article-api/internal/di"
)

func main() {
	cfg := di.InitializedViper()
	logger := di.InitializedLogrus(cfg)

	config.MigrateDB(cfg, logger)

	server := di.InitializedServer(cfg, logger)
	server.SetupRoutes()
	server.StartServer()
}
