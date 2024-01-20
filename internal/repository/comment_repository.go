package repository

import (
	"github.com/jordanmarcelino/go-article-api/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CommentRepository struct {
	Repository[entity.Comment]
	Log *logrus.Logger
}

func NewCommentRepository(log *logrus.Logger) *CommentRepository {
	return &CommentRepository{Log: log}
}

func (r *CommentRepository) FindUserComments(db *gorm.DB, id string) ([]entity.Comment, error) {
	var comments []entity.Comment
	if err := db.Where("user_id = ?", id).Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *CommentRepository) FindByIdWithRelation(db *gorm.DB, comment *entity.Comment, id string) error {

	return db.Preload("User").Preload("Article").Take(comment, "id = ?", id).Error
}

func (r *CommentRepository) DeleteAllByArticleId(db *gorm.DB, id string) error {

	return db.Where("article_id = ?", id).Delete(&entity.Comment{}).Error
}
