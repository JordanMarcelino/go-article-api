package repository

import (
	"github.com/jordanmarcelino/go-article-api/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ArticleRepository struct {
	Repository[entity.Article]
	Log *logrus.Logger
}

func NewArticleRepository(log *logrus.Logger) *ArticleRepository {
	return &ArticleRepository{Log: log}
}

func (r *ArticleRepository) FindAllWithRelation(db *gorm.DB) ([]entity.Article, error) {
	var articles []entity.Article

	if err := db.Preload("Tags").Preload("Comments").Find(&articles).Error; err != nil {
		return nil, err
	}

	return articles, nil
}

func (r *ArticleRepository) FindByIdWithRelation(db *gorm.DB, article *entity.Article, id string) error {
	return db.Preload("Tags").Preload("Comments").Take(article, "id = ?",
		id).Error
}

func (r *ArticleRepository) CreateWithRelation(db *gorm.DB, article *entity.Article) error {
	return db.Create(article).Association("Tags").Append(&article.Tags)
}

func (r *ArticleRepository) UpdateWithRelation(db *gorm.DB, article *entity.Article) error {
	return db.Model(article).Association("Tags").Replace(&article.Tags)
}

func (r *ArticleRepository) DeleteWithRelation(db *gorm.DB, article *entity.Article) error {
	return db.Model(article).Association("Tags").Clear()
}
