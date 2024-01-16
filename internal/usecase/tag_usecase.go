package usecase

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jordanmarcelino/go-article-api/internal/entity"
	"github.com/jordanmarcelino/go-article-api/internal/model"
	"github.com/jordanmarcelino/go-article-api/internal/model/converter"
	"github.com/jordanmarcelino/go-article-api/internal/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TagUseCase struct {
	DB            *gorm.DB
	Log           *logrus.Logger
	Validate      *validator.Validate
	TagRepository *repository.TagRepository
}

func NewTagUseCase(DB *gorm.DB, log *logrus.Logger, validate *validator.Validate, tagRepository *repository.TagRepository) *TagUseCase {
	return &TagUseCase{DB: DB, Log: log, Validate: validate, TagRepository: tagRepository}
}

func (c *TagUseCase) Create(ctx context.Context, request *model.CreateTagRequest) (*model.TagResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.StructCtx(ctx, request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	tag := new(entity.Tag)
	if err := c.TagRepository.Create(tx, tag); err != nil {
		c.Log.Warnf("Failed create tag to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.TagToResponse(tag), nil
}

func (c *TagUseCase) Update(ctx context.Context, request *model.UpdateTagRequest) (*model.TagResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.StructCtx(ctx, request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	tag := new(entity.Tag)
	tag.ID = request.Id
	tag.Name = request.Name

	if err := c.TagRepository.Update(tx, tag); err != nil {
		c.Log.Warnf("Failed update tag to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.TagToResponse(tag), nil
}

func (c *TagUseCase) Delete(ctx context.Context, request *model.DeleteTagRequest) (*model.TagResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.StructCtx(ctx, request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	tag := new(entity.Tag)
	tag.ID = request.Id

	if err := c.TagRepository.Delete(tx, tag); err != nil {
		c.Log.Warnf("Failed delete tag to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.TagToResponse(tag), nil
}

func (c *TagUseCase) Get(ctx context.Context, request *model.GetTagRequest) (*model.TagResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.StructCtx(ctx, request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	tag := new(entity.Tag)

	if err := c.TagRepository.FindById(tx, tag, request.Id); err != nil {
		c.Log.Warnf("Failed create tag to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.TagToResponse(tag), nil
}

func (c *TagUseCase) List(ctx context.Context) ([]model.TagResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	tags, err := c.TagRepository.FindAll(tx)
	if err != nil {
		c.Log.Warnf("Failed update tag to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	var responses []model.TagResponse
	for _, tag := range tags {
		responses = append(responses, *converter.TagToResponse(&tag))
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return responses, nil
}
