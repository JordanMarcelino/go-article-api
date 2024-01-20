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

type ArticleUseCase struct {
	DB                *gorm.DB
	Log               *logrus.Logger
	Validate          *validator.Validate
	ArticleRepository *repository.ArticleRepository
	CommentRepository *repository.CommentRepository
	TagRepository     *repository.TagRepository
}

func NewArticleUseCase(DB *gorm.DB, log *logrus.Logger, validate *validator.Validate, articleRepository *repository.ArticleRepository) *ArticleUseCase {
	return &ArticleUseCase{DB: DB, Log: log, Validate: validate, ArticleRepository: articleRepository}
}

func (c *ArticleUseCase) Create(ctx context.Context, request *model.CreateArticleRequest) (*model.ArticleResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.StructCtx(ctx, request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	var tagIds []int64

	for _, tag := range request.Tags {
		tagIds = append(tagIds, tag.Id)
	}

	tags, err := c.TagRepository.FindByIds(tx, tagIds)
	if err != nil {
		c.Log.Warnf("Failed to get tags from database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	article := &entity.Article{
		ID:          uuid.NewString(),
		Thumbnail:   request.Thumbnail,
		Title:       request.Title,
		Description: request.Description,
		Body:        request.Body,
		UserId:      request.UserId,
		Tags:        tags,
	}

	if err := c.ArticleRepository.CreateWithRelation(tx, article); err != nil {
		c.Log.Warnf("Failed create article to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.ArticleToResponse(article), nil
}

func (c *ArticleUseCase) Update(ctx context.Context, request *model.UpdateArticleRequest) (*model.ArticleResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.StructCtx(ctx, request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	article := new(entity.Article)
	if err := c.ArticleRepository.FindByIdWithRelation(tx, article, request.Id); err != nil {
		c.Log.Warnf("Failed to get article from database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if len(request.Tags) > 0 {
		var tagIds []int64

		for _, tag := range request.Tags {
			tagIds = append(tagIds, tag.Id)
		}

		tags, err := c.TagRepository.FindByIds(tx, tagIds)
		if err != nil {
			c.Log.Warnf("Failed to get tags from database : %+v", err)
			return nil, fiber.ErrInternalServerError
		}

		article.Tags = tags
	}

	if request.Thumbnail != "" {
		article.Thumbnail = request.Thumbnail
	}

	if request.Title != "" {
		article.Title = request.Title
	}

	if request.Description != "" {
		article.Description = request.Description
	}

	if request.Body != "" {
		article.Body = request.Body
	}

	if err := c.ArticleRepository.Update(tx, article); err != nil {
		c.Log.Warnf("Failed update article to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := c.ArticleRepository.UpdateWithRelation(tx, article); err != nil {
		c.Log.Warnf("Failed update article tags to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.ArticleToResponse(article), nil
}

func (c *ArticleUseCase) Get(ctx context.Context, request *model.GetArticleRequest) (*model.ArticleResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.StructCtx(ctx, request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	article := new(entity.Article)

	if err := c.ArticleRepository.FindByIdWithRelation(tx, article, request.Id); err != nil {
		c.Log.Warnf("Failed get article from database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.ArticleToResponse(article), nil
}

func (c *ArticleUseCase) Delete(ctx context.Context, request *model.DeleteArticleRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.StructCtx(ctx, request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return fiber.ErrBadRequest
	}

	article := new(entity.Article)
	if err := c.ArticleRepository.FindByIdWithRelation(tx, article, request.Id); err != nil {
		c.Log.Warnf("Failed get article from database : %+v", err)
		return fiber.ErrInternalServerError
	}

	if err := c.ArticleRepository.DeleteWithRelation(tx, article); err != nil {
		c.Log.Warnf("Failed delete article tags from database : %+v", err)
		return fiber.ErrInternalServerError
	}

	if err := c.CommentRepository.DeleteAllByArticleId(tx, request.Id); err != nil {
		c.Log.Warnf("Failed delete article comments from database : %+v", err)
		return fiber.ErrInternalServerError
	}

	if err := c.ArticleRepository.Delete(tx, article); err != nil {
		c.Log.Warnf("Failed delete article from database : %+v", err)
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return fiber.ErrInternalServerError
	}

	return nil
}

func (c *ArticleUseCase) List(ctx context.Context) ([]model.ArticleResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	articles, err := c.ArticleRepository.FindAllWithRelation(tx)
	if err != nil {
		c.Log.Warnf("Failed get article from database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.ArticlesToResponses(articles), nil
}
