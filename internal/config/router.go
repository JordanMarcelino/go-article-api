package config

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jordanmarcelino/go-article-api/internal/controller"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type RouteConfig struct {
	App               *fiber.App
	Config            *viper.Viper
	Log               *logrus.Logger
	UserController    *controller.UserController
	ArticleController *controller.ArticleController
	TagController     *controller.TagController
	CommentController *controller.CommentController
	AuthMiddleware    fiber.Handler
}

func NewRouteConfig(app *fiber.App, config *viper.Viper, log *logrus.Logger, userController *controller.UserController, articleController *controller.ArticleController, tagController *controller.TagController, commentController *controller.CommentController, authMiddleware fiber.Handler) *RouteConfig {
	return &RouteConfig{App: app, Config: config, Log: log, UserController: userController, ArticleController: articleController, TagController: tagController, CommentController: commentController, AuthMiddleware: authMiddleware}
}

func SetupRoutes(c *RouteConfig) {
	api := c.App.Group("/api")
	userApi := api.Group("/users")
	articleApi := api.Group("/articles")
	tagApi := api.Group("/tags")
	commentApi := api.Group("/comments")

	userApi.Post("/signup", c.UserController.Register)
	userApi.Post("/login", c.UserController.Login)
	userApi.Get("/:userId", c.AuthMiddleware, c.UserController.Update)
	userApi.Put("/:userId", c.AuthMiddleware, c.UserController.Update)
	userApi.Put("/:userId/comments", c.AuthMiddleware, c.CommentController.User)

	articleApi.Use(c.AuthMiddleware)
	articleApi.Get("", c.ArticleController.List)
	articleApi.Post("", c.ArticleController.Create)
	articleApi.Get("/:articleId", c.ArticleController.Get)
	articleApi.Put("/:articleId", c.ArticleController.Update)
	articleApi.Delete("/:articleId", c.ArticleController.Delete)

	tagApi.Use(c.AuthMiddleware)
	tagApi.Get("", c.TagController.List)
	tagApi.Post("", c.TagController.Create)
	tagApi.Get("/:tagId", c.TagController.Get)
	tagApi.Put("/:tagId", c.TagController.Update)
	tagApi.Delete("/:tagId", c.TagController.Delete)

	commentApi.Use(c.AuthMiddleware)
	commentApi.Post("", c.CommentController.Create)
	commentApi.Get("/:commentId", c.CommentController.Get)
	commentApi.Delete("/:commentId", c.CommentController.Delete)
}

func StartServer(c *RouteConfig) {
	webPort := c.Config.GetInt("web.port")
	err := c.App.Listen(fmt.Sprintf(":%d", webPort))

	if err != nil {
		c.Log.Fatalf("Failed to start server : %+v", err)
	}
}
