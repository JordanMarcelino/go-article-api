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
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CommentUseCase struct {
	DB                *gorm.DB
	Log               *logrus.Logger
	Validate          *validator.Validate
	CommentRepository *repository.CommentRepository
	UseRepository     *repository.UserRepository
}

func NewCommentUseCase(DB *gorm.DB, log *logrus.Logger, validate *validator.Validate, commentRepository *repository.CommentRepository) *CommentUseCase {
	return &CommentUseCase{DB: DB, Log: log, Validate: validate, CommentRepository: commentRepository}
}

func (c *CommentUseCase) Create(ctx context.Context, request *model.CreateCommentRequest) (*model.CommentResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.StructCtx(ctx, request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	comment := new(entity.Comment)
	comment.ID = uuid.NewString()
	comment.Body = request.Body
	comment.UserId = request.UserId
	comment.ArticleId = request.ArticleId

	if err := c.CommentRepository.Create(tx, comment); err != nil {
		c.Log.Warnf("Failed create comment to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.CommentToResponse(comment), nil
}

func (c *CommentUseCase) Get(ctx context.Context, request *model.GetCommentRequest) (*model.CommentResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.StructCtx(ctx, request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	comment := new(entity.Comment)

	if err := c.CommentRepository.FindById(tx, comment, request.Id); err != nil {
		c.Log.Warnf("Failed get comment from database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.CommentToResponse(comment), nil
}

func (c *CommentUseCase) Delete(ctx context.Context, request *model.DeleteCommentRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.StructCtx(ctx, request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return fiber.ErrBadRequest
	}

	comment := new(entity.Comment)
	comment.ID = request.Id

	if err := c.CommentRepository.Delete(tx, comment); err != nil {
		c.Log.Warnf("Failed get comment from database : %+v", err)
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return fiber.ErrInternalServerError
	}

	return nil
}

func (c *CommentUseCase) User(ctx context.Context, request *model.GetUserCommentRequest) ([]model.CommentResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.StructCtx(ctx, request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	comments, err := c.CommentRepository.FindUserComments(tx, request.Id)

	if err != nil {
		c.Log.Warnf("Failed get user comment from database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.CommentsToResponse(comments), nil
}
