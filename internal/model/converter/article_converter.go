package converter

import (
	"github.com/jordanmarcelino/go-article-api/internal/entity"
	"github.com/jordanmarcelino/go-article-api/internal/model"
)

func ArticleToResponse(article *entity.Article) *model.ArticleResponse {
	return &model.ArticleResponse{
		Id:          article.ID,
		Thumbnail:   article.Thumbnail,
		Title:       article.Title,
		Description: article.Description,
		Body:        article.Body,
		CreatedAt:   article.CreatedAt,
		UpdatedAt:   article.UpdatedAt,
		UserId:      article.UserId,
		Tags:        TagsToResponse(article.Tags),
		Comments:    CommentsToResponse(article.Comments),
	}
}
