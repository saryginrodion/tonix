package jwt

type JWTError struct {
	Message string `json:"message"`
}

func NewJWTError(message string) *JWTError {
	return &JWTError{
		Message: message,
	}
}

func (e *JWTError) Error() string {
	return e.Message
}
