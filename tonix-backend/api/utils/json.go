package utils

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator/v10"
)


func ParseAndValidateJson[T any](b io.ReadCloser, res T) (*T, error) {
	val, err := io.ReadAll(b)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(val, &res)

	if err != nil {
		return nil, err
	}

	validate := validator.New()

	err = validate.Struct(res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
