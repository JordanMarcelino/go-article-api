package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jordanmarcelino/go-article-api/internal/model"
	"github.com/spf13/viper"
)

func NewFiber(viper *viper.Viper) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      viper.GetString("app.name"),
		ErrorHandler: NewErrorHandler(),
		Prefork:      viper.GetBool("web.prefork"),
	})

	return app
}

func NewErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		return ctx.Status(code).JSON(model.WebResponse[any]{
			Info: map[string]any{
				"success": false,
				"meta": map[string]any{
					"errors": err.Error(),
				},
				"message": "Internal server error",
			},
			Data: nil,
		})
	}
}
