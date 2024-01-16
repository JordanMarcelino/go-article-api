package repository

import (
	"github.com/jordanmarcelino/go-article-api/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	Repository[entity.User]
	Log *logrus.Logger
}

func NewUserRepository(log *logrus.Logger) *UserRepository {
	return &UserRepository{Log: log}
}

func (r *UserRepository) CountByEmail(db *gorm.DB, email string) (int64, error) {
	var total int64
	err := db.Model(new(entity.User)).Where("email = ?", email).Count(&total).Error

	return total, err
}

func (r *UserRepository) FindByEmail(db *gorm.DB, user *entity.User, email string) error {
	return db.Where("email = ?", email).Take(user).Error
}
