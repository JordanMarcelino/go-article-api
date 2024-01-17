package repository

import (
	"github.com/jordanmarcelino/go-article-api/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TagRepository struct {
	Repository[entity.Tag]
	Log *logrus.Logger
}

func NewTagRepository(log *logrus.Logger) *TagRepository {
	return &TagRepository{Log: log}
}

func (r *TagRepository) FindByIds(db *gorm.DB, ids []int64) ([]entity.Tag, error) {
	var tags []entity.Tag

	if err := db.Find(&tags, "id IN ?", ids).Error; err != nil {
		return nil, err
	}

	return tags, nil
}
