package model

type TagResponse struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type CreateTagRequest struct {
	Name string `json:"name" validate:"required,min=3,max=100"`
}

type UpdateTagRequest struct {
	Id   int64  `json:"id" validate:"required"`
	Name string `json:"name" validate:"required,min=3,max=100"`
}

type GetTagRequest struct {
	Id int64 `json:"id" validate:"required"`
}

type DeleteTagRequest struct {
	Id int64 `json:"id" validate:"required"`
}
