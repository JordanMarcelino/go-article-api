package repository

import (
	"github.com/jordanmarcelino/go-article-api/internal/entity"
	"github.com/sirupsen/logrus"
)

type TagRepository struct {
	Repository[entity.Tag]
	Log *logrus.Logger
}

func NewTagRepository(log *logrus.Logger) *TagRepository {
	return &TagRepository{Log: log}
}
