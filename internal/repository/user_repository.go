package repository

import (
	"github.com/jordanmarcelino/go-article-api/internal/entity"
	"github.com/sirupsen/logrus"
)

type UserRepository struct {
	Repository[entity.User]
	Log *logrus.Logger
}

func NewUserRepository(log *logrus.Logger) *UserRepository {
	return &UserRepository{Log: log}

}
