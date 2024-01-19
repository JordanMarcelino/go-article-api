package usecase

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jordanmarcelino/go-article-api/internal/entity"
	"github.com/jordanmarcelino/go-article-api/internal/model"
	"github.com/jordanmarcelino/go-article-api/internal/model/converter"
	"github.com/jordanmarcelino/go-article-api/internal/repository"
	"github.com/jordanmarcelino/go-article-api/internal/util"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserUseCase struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	Validate       *validator.Validate
	UserRepository *repository.UserRepository
}

func NewUserUseCase(DB *gorm.DB, log *logrus.Logger, validate *validator.Validate, userRepository *repository.UserRepository) *UserUseCase {
	return &UserUseCase{DB: DB, Log: log, Validate: validate, UserRepository: userRepository}
}

func (c *UserUseCase) Login(ctx context.Context, request *model.LoginUserRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.StructCtx(ctx, request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return fiber.ErrBadRequest
	}

	total, err := c.UserRepository.CountByEmail(tx, request.Email)

	if err != nil {
		c.Log.Warnf("Failed to count user from database : %+v", err)
		return fiber.ErrInternalServerError
	}

	if total == 0 {
		c.Log.Warnf("User doesn't exist : %+v", err)
		return fiber.ErrNotFound
	}

	user := new(entity.User)
	if err := c.UserRepository.FindByEmail(tx, user, request.Email); err != nil {
		c.Log.Warnf("Failed to get user : %+v", err)
		return fiber.ErrInternalServerError
	}

	if authorize := util.CheckPasswordHash(request.Password, user.Password); !authorize {
		c.Log.Warnf("Password doesn't match")
		return fiber.ErrUnauthorized
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return fiber.ErrInternalServerError
	}

	return nil
}

func (c *UserUseCase) Create(ctx context.Context, request *model.RegisterUserRequest) (*model.UserResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.StructCtx(ctx, request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	total, err := c.UserRepository.CountByEmail(tx, request.Email)

	if err != nil {
		c.Log.Warnf("Failed to count user from database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if total > 0 {
		c.Log.Warnf("User already exist : %+v", err)
		return nil, fiber.ErrConflict
	}

	password, err := util.HashPassword(request.Password)
	if err != nil {
		c.Log.Warnf("Failed to generate bcrypt hash : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	id := uuid.NewString()
	user := &entity.User{
		ID:       id,
		Username: "User " + id,
		Email:    request.Email,
		Password: password,
		Avatar:   "https://avatars.githubusercontent.com/u/42530587?v=4",
	}

	if err := c.UserRepository.Create(tx, user); err != nil {
		c.Log.Warnf("Failed create user to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.UserToResponse(user), nil
}

func (c *UserUseCase) Update(ctx context.Context, request *model.UpdateUserRequest) (*model.UserResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.StructCtx(ctx, request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	total, err := c.UserRepository.CountById(tx, request.Id)

	if err != nil {
		c.Log.Warnf("Failed to count user from database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if total == 0 {
		c.Log.Warnf("User doesn't exist : %+v", err)
		return nil, fiber.ErrNotFound
	}

	user := new(entity.User)

	c.UserRepository.FindById(tx, user, request.Id)

	if request.Password != "" {
		password, err := util.HashPassword(request.Password)
		if err != nil {
			c.Log.Warnf("Failed to generate bcrypt hash : %+v", err)
			return nil, fiber.ErrInternalServerError
		}
		user.Password = password
	}

	if request.Username != "" {
		user.Username = request.Username
	}

	if request.Email != "" {
		user.Email = request.Email
	}

	if request.Phone != "" {
		user.Phone = request.Phone
	}

	if request.Avatar != "" {
		user.Avatar = request.Avatar
	}

	if err := c.UserRepository.Update(tx, user); err != nil {
		c.Log.Warnf("Failed update user to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.UserToResponse(user), nil
}

func (c *UserUseCase) FindById(ctx context.Context, id string) (*model.UserResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.VarCtx(ctx, id, "uuid4"); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	user := new(entity.User)
	total, err := c.UserRepository.CountById(tx, id)

	if err != nil {
		c.Log.Warnf("Failed to count user from database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if total == 0 {
		c.Log.Warnf("User doesn't exist : %+v", err)
		return nil, fiber.ErrNotFound
	}

	if err := c.UserRepository.FindById(tx, user, id); err != nil {
		c.Log.Warnf("Failed find user from database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.UserToResponse(user), nil
}
