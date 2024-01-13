package model

type ArticleResponse struct {
	Id          string            `json:"id"`
	Thumbnail   string            `json:"thumbnail"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Body        string            `json:"body"`
	CreatedAt   int64             `json:"created_at"`
	UpdatedAt   int64             `json:"updated_at"`
	UserId      string            `json:"user_id"`
	Tags        []TagResponse     `json:"tags"`
	Comments    []CommentResponse `json:"comments"`
}

type CreateArticleRequest struct {
	Thumbnail   string `json:"thumbnail" validate:"required"`
	Title       string `json:"title" validate:"required,max=100"`
	Description string `json:"description" validate:"required,max=100"`
	Body        string `json:"body" validate:"required,min=1"`
	UserId      string `json:"user_id" validate:"required,uuid4"`
}

type UpdateArticleRequest struct {
	Id          string `json:"id" validate:"required,uuid4"`
	Thumbnail   string `json:"thumbnail"`
	Title       string `json:"title" validate:"max=100"`
	Description string `json:"description" validate:"max=100"`
	Body        string `json:"body" validate:"min=1"`
}

type DeleteArticleRequest struct {
	Id string `json:"id" validate:"required,uuid4"`
}

type GetArticleRequest struct {
	Id string `json:"id" validate:"required,uuid4"`
}
