package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jordanmarcelino/go-article-api/internal/model"
	"github.com/jordanmarcelino/go-article-api/internal/usecase"
	"github.com/sirupsen/logrus"
)

type CommentController struct {
	Log     *logrus.Logger
	UseCase *usecase.CommentUseCase
}

func NewCommentController(log *logrus.Logger, useCase *usecase.CommentUseCase) *CommentController {
	return &CommentController{Log: log, UseCase: useCase}
}

func (c *CommentController) Create(ctx *fiber.Ctx) error {
	request := new(model.CreateCommentRequest)

	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to create comment : %+v", err)
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(model.WebResponse[*model.CommentResponse]{
		Info: map[string]any{
			"success": true,
			"meta":    nil,
			"message": "Success create new comment",
		},
		Data: response,
	})
}

func (c *CommentController) Get(ctx *fiber.Ctx) error {
	request := new(model.GetCommentRequest)
	request.Id = ctx.Params("commentId")

	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to get comment : %+v", err)
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(model.WebResponse[*model.CommentResponse]{
		Info: map[string]any{
			"success": true,
			"meta":    nil,
			"message": "Success get comment",
		},
		Data: response,
	})
}

func (c *CommentController) User(ctx *fiber.Ctx) error {
	request := new(model.GetUserCommentRequest)
	request.Id = ctx.Params("userID")

	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.User(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to get comments : %+v", err)
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(model.WebResponse[[]model.CommentResponse]{
		Info: map[string]any{
			"success": true,
			"meta":    nil,
			"message": "Success get user comments",
		},
		Data: response,
	})
}

func (c *CommentController) Delete(ctx *fiber.Ctx) error {
	request := new(model.DeleteCommentRequest)
	request.Id = ctx.Params("commentId")

	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	err := c.UseCase.Delete(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to delete comment : %+v", err)
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(model.WebResponse[*model.CommentResponse]{
		Info: map[string]any{
			"success": true,
			"meta":    nil,
			"message": "Success delete comment",
		},
		Data: nil,
	})
}
