package usersdto

type CreateUserRequest struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Phone    int    `json:"phone" form:"phone" validate:"required"`
	Address  string `json:"address" form:"address" validate:"required"`
	Gender   string `json:"gender" form:"gender" validate:"required"`
	Image    string `json:"image" form:"image"`
	Password string `json:"password" form:"password" validate:"required"`
}

type UpdateUserRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    int    `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
	Gender   string `json:"gender" form:"gender"`
	Image    string `json:"image" form:"image"`
}
