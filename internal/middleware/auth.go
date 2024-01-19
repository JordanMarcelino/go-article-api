package middleware

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/jordanmarcelino/go-article-api/internal/config"
	"github.com/jordanmarcelino/go-article-api/internal/model"
)

// Protected protect routes
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(config.GetEnv("JWT_SECRET", "secret"))},
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(model.WebResponse[map[string]any]{
				Info: map[string]any{
					"success": false,
					"meta":    nil,
					"message": "Missing or malformed JWT",
				},
				Data: nil,
			})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(model.WebResponse[map[string]any]{
			Info: map[string]any{
				"success": false,
				"meta":    nil,
				"message": "Invalid or expired JWT",
			},
			Data: nil,
		})
}
