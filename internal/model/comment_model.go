package model

type CommentResponse struct {
	Id        string `json:"id"`
	Body      string `json:"body"`
	UserId    string `json:"user_id"`
	ArticleId string `json:"article_id"`
	CreatedAt int64  `json:"created_at"`
}

type CreateCommentRequest struct {
	UserId    string `json:"user_id" validate:"required,uuid4"`
	ArticleId string `json:"article_id" validate:"required,uuid4"`
	Body      string `json:"body" validate:"required,min=1,max=255"`
}
