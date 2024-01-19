package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jordanmarcelino/go-article-api/internal/model"
	"github.com/jordanmarcelino/go-article-api/internal/usecase"
	"github.com/sirupsen/logrus"
	"strconv"
)

type TagController struct {
	Log     *logrus.Logger
	UseCase *usecase.TagUseCase
}

func NewTagController(log *logrus.Logger, useCase *usecase.TagUseCase) *TagController {
	return &TagController{Log: log, UseCase: useCase}
}

func (c *TagController) Create(ctx *fiber.Ctx) error {
	request := new(model.CreateTagRequest)

	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to create tag : %+v", err)
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(model.WebResponse[*model.TagResponse]{
		Info: map[string]any{
			"success": true,
			"meta":    nil,
			"message": "Success create tag",
		},
		Data: response,
	})
}

func (c *TagController) Update(ctx *fiber.Ctx) error {
	request := new(model.UpdateTagRequest)
	tagId := ctx.Params("tagId")

	id, err := strconv.Atoi(tagId)
	if err != nil {
		c.Log.Warnf("Invalid id params : %+v", err)
		return fiber.ErrBadRequest
	}

	request.Id = int64(id)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to update tag : %+v", err)
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(model.WebResponse[*model.TagResponse]{
		Info: map[string]any{
			"success": true,
			"meta":    nil,
			"message": "Success update existing tag",
		},
		Data: response,
	})
}

func (c *TagController) Delete(ctx *fiber.Ctx) error {
	request := new(model.DeleteTagRequest)
	tagId := ctx.Params("tagId")

	id, err := strconv.Atoi(tagId)
	if err != nil {
		c.Log.Warnf("Invalid id params : %+v", err)
		return fiber.ErrBadRequest
	}

	request.Id = int64(id)

	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	err = c.UseCase.Delete(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to delete tag : %+v", err)
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(model.WebResponse[*model.TagResponse]{
		Info: map[string]any{
			"success": true,
			"meta":    nil,
			"message": "Success delete tag",
		},
		Data: nil,
	})
}

func (c *TagController) Get(ctx *fiber.Ctx) error {
	request := new(model.GetTagRequest)
	tagId := ctx.Params("tagId")

	id, err := strconv.Atoi(tagId)
	if err != nil {
		c.Log.Warnf("Invalid id params : %+v", err)
		return fiber.ErrBadRequest
	}

	request.Id = int64(id)

	response, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to get tag : %+v", err)
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(model.WebResponse[*model.TagResponse]{
		Info: map[string]any{
			"success": true,
			"meta":    nil,
			"message": "Success get tag",
		},
		Data: response,
	})
}

func (c *TagController) List(ctx *fiber.Ctx) error {
	response, err := c.UseCase.List(ctx.UserContext())
	if err != nil {
		c.Log.Warnf("Failed to get all tag : %+v", err)
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(model.WebResponse[[]model.TagResponse]{
		Info: map[string]any{
			"success": true,
			"meta":    nil,
			"message": "Success get all tag",
		},
		Data: response,
	})
}
