package requests

type RegistrationBody struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
	Username string `json:"username" validate:"required,username,min=3,max=32"`
}

type LoginBody struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}
