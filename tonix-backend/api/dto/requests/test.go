package requests

type TestMessage struct {
	Message string `json:"message" validate:"required,email"`
}
