package converter

import (
	"github.com/jordanmarcelino/go-article-api/internal/entity"
	"github.com/jordanmarcelino/go-article-api/internal/model"
)

func CommentToResponse(comment *entity.Comment) *model.CommentResponse {
	return &model.CommentResponse{
		Id:        comment.ID,
		Body:      comment.Body,
		UserId:    comment.User.ID,
		ArticleId: comment.Article.ID,
		CreatedAt: comment.CreatedAt,
	}
}

func CommentsToResponse(comments []entity.Comment) []*model.CommentResponse {
	var commentResponses []*model.CommentResponse

	for _, comment := range comments {
		commentResponses = append(commentResponses, CommentToResponse(&comment))
	}

	return commentResponses
}
