package response_wrapper

type APIErrorResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
	Errors  []any  `json:"errors"`
}

func ErrorsResponse(message string, errors ...any) APIErrorResponse {
	return APIErrorResponse{
		Ok:     false,
		Message: message,
		Errors: errors,
	}
}
