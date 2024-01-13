package model

type CommentResponse struct {
	Id        string `json:"id"`
	Body      string `json:"body"`
	CreatedAt int64  `json:"created_at"`
}

type CreateCommentRequest struct {
	Body string `json:"body" validate:"required,min=1,max=255"`
}
