package response_wrapper

type APIResponse[T any] struct {
	Ok   bool `json:"ok"`
	Data T    `json:"data"`
}

func OkResponse[T any](data T) APIResponse[T] {
	return APIResponse[T]{
		Ok:   true,
		Data: data,
	}
}
