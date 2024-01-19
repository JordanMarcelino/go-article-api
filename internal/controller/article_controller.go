package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jordanmarcelino/go-article-api/internal/model"
	"github.com/jordanmarcelino/go-article-api/internal/usecase"
	"github.com/sirupsen/logrus"
)

type ArticleController struct {
	Log     *logrus.Logger
	UseCase *usecase.ArticleUseCase
}

func NewArticleController(log *logrus.Logger, useCase *usecase.ArticleUseCase) *ArticleController {
	return &ArticleController{Log: log, UseCase: useCase}
}

func (c *ArticleController) Create(ctx *fiber.Ctx) error {
	request := new(model.CreateArticleRequest)

	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to create article : %+v", err)
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(model.WebResponse[*model.ArticleResponse]{
		Info: map[string]any{
			"success": true,
			"meta":    nil,
			"message": "Success create new article",
		},
		Data: response,
	})
}

func (c *ArticleController) Update(ctx *fiber.Ctx) error {
	request := new(model.UpdateArticleRequest)
	request.Id = ctx.Params("articleId")

	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to update article : %+v", err)
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(model.WebResponse[*model.ArticleResponse]{
		Info: map[string]any{
			"success": true,
			"meta":    nil,
			"message": "Success update existing article",
		},
		Data: response,
	})
}

func (c *ArticleController) Delete(ctx *fiber.Ctx) error {
	request := new(model.DeleteArticleRequest)
	request.Id = ctx.Params("articleId")

	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	err := c.UseCase.Delete(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to delete article : %+v", err)
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(model.WebResponse[*model.ArticleResponse]{
		Info: map[string]any{
			"success": true,
			"meta":    nil,
			"message": "Success delete article",
		},
		Data: nil,
	})
}

func (c *ArticleController) Get(ctx *fiber.Ctx) error {
	request := new(model.GetArticleRequest)
	request.Id = ctx.Params("articleId")

	response, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to get article : %+v", err)
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(model.WebResponse[*model.ArticleResponse]{
		Info: map[string]any{
			"success": true,
			"meta":    nil,
			"message": "Success get article",
		},
		Data: response,
	})
}

func (c *ArticleController) List(ctx *fiber.Ctx) error {
	response, err := c.UseCase.List(ctx.UserContext())
	if err != nil {
		c.Log.Warnf("Failed to get all article : %+v", err)
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(model.WebResponse[[]model.ArticleResponse]{
		Info: map[string]any{
			"success": true,
			"meta":    nil,
			"message": "Success get all article",
		},
		Data: response,
	})
}
