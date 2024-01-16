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
	if err := db.Preload("User", "id = ?", id).Preload("Article").Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *CommentRepository) FindByIdWithRelation(db *gorm.DB, comment *entity.Comment, id string) error {

	return db.Preload("User").Preload("Article").Take(comment, "id = ?", id).Error
}
