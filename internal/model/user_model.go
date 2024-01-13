package model

type UserResponse struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type RegisterUserRequest struct {
	Email    string `json:"email" validate:"required,email,max=100"`
	Password string `json:"password" validate:"required,min=5,max=100"`
}

type UpdateUserRequest struct {
	Id       string `json:"id" validate:"required,uuid4"`
	Username string `json:"username" validate:"max=100"`
	Password string `json:"password"`
	Email    string `json:"email" validate:"email,max=100"`
	Phone    string `json:"phone" validate:"e164"`
	Avatar   string `json:"avatar" validate:"max=100"`
}

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
