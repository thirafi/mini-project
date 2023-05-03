package payload

type CreateUserRequest struct {
	Name     string `json:"name" form:"name" validate:"required,max=20"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=5"`
}
