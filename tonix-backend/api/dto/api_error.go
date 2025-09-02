package dto

type ApiError struct {
	Status int
	Message string
}

func NewApiError(status int, message string) *ApiError {
	return &ApiError{
		status,
		message,
	}
}

func (e *ApiError) Error() string {
	return e.Message
}
