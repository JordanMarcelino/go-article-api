package model

type TagResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CreateTagRequest struct {
	Name string `json:"name" validate:"required,min=3,max=100"`
}

type UpdateTagRequest struct {
	Id   string `json:"id" validate:"required,uuid4"`
	Name string `json:"name" validate:"required,min=3,max=100"`
}

type GetTagRequest struct {
	Id string `json:"id" validate:"required,uuid4"`
}

type DeleteTagRequest struct {
	Id string `json:"id" validate:"required,uuid4"`
}
