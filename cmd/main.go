package main

import (
	"github.com/jordanmarcelino/go-article-api/internal/config"
	"github.com/jordanmarcelino/go-article-api/internal/di"
)

func main() {
	app := di.InitializedServer()

	config.SetupRoutes(app)
	config.StartServer(app)
}
