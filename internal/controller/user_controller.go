package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jordanmarcelino/go-article-api/internal/model"
	"github.com/jordanmarcelino/go-article-api/internal/usecase"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

type UserController struct {
	Log     *logrus.Logger
	Config  *viper.Viper
	UseCase *usecase.UserUseCase
}

func NewUserController(log *logrus.Logger, useCase *usecase.UserUseCase) *UserController {
	return &UserController{Log: log, UseCase: useCase}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	request := new(model.RegisterUserRequest)

	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to register user : %+v", err)
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(model.WebResponse[*model.UserResponse]{
		Info: map[string]any{
			"success": true,
			"meta":    nil,
			"message": "Success register new user",
		},
		Data: response,
	})
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	request := new(model.LoginUserRequest)

	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	err := c.UseCase.Login(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to login : %+v", err)
		return fiber.ErrInternalServerError
	}

	claims := jwt.MapClaims{
		"email": request.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(c.Config.GetString("jwt.secret")))
	if err != nil {
		c.Log.Warnf("Failed to generate jwt token : %+v", err)
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(model.WebResponse[map[string]any]{
		Info: map[string]any{
			"success": true,
			"meta":    nil,
			"message": "Success login",
		},
		Data: map[string]any{
			"token": t,
		},
	})
}

func (c *UserController) Update(ctx *fiber.Ctx) error {
	request := new(model.UpdateUserRequest)

	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	userId := ctx.Params("userId")
	request.Id = userId

	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to update user : %+v", err)
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(model.WebResponse[*model.UserResponse]{
		Info: map[string]any{
			"success": true,
			"meta":    nil,
			"message": "Success update existing user",
		},
		Data: response,
	})
}
