//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/jordanmarcelino/go-article-api/internal/config"
	"github.com/jordanmarcelino/go-article-api/internal/controller"
	"github.com/jordanmarcelino/go-article-api/internal/middleware"
	"github.com/jordanmarcelino/go-article-api/internal/repository"
	"github.com/jordanmarcelino/go-article-api/internal/usecase"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var userSet = wire.NewSet(
	repository.NewUserRepository,
	usecase.NewUserUseCase,
	controller.NewUserController,
)

var articleSet = wire.NewSet(
	repository.NewArticleRepository,
	usecase.NewArticleUseCase,
	controller.NewArticleController,
)

var tagSet = wire.NewSet(
	repository.NewTagRepository,
	usecase.NewTagUseCase,
	controller.NewTagController,
)

var commentSet = wire.NewSet(
	repository.NewCommentRepository,
	usecase.NewCommentUseCase,
	controller.NewCommentController,
)

func InitializedViper() *viper.Viper {
	wire.Build(config.NewViper)
	return nil
}

func InitializedLogrus(cfg *viper.Viper) *logrus.Logger {
	wire.Build(config.NewLogger)
	return nil
}

func InitializedServer(cfg *viper.Viper, logger *logrus.Logger) *config.RouteConfig {
	wire.Build(
		config.NewFiber,
		config.NewValidator,
		config.NewDatabase,
		userSet,
		articleSet,
		tagSet,
		commentSet,
		middleware.Protected,
		config.NewRouteConfig,
	)

	return nil
}
