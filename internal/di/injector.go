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

func InitializedServer() *config.RouteConfig {
	wire.Build(
		config.NewViper,
		config.NewLogger,
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
