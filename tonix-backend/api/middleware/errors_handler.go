package middleware

import (
	"encoding/json"
	"errors"
	"net/http"
	"tonix/backend/api/context"
	"tonix/backend/api/dto"
	wrap "tonix/backend/api/dto/response_wrapper"

	"github.com/go-playground/validator/v10"
	"github.com/saryginrodion/stackable"
)

var ErrorsHandlerMiddleware = stackable.WrapFunc(
	func(ctx *context.Context, next func() error) error {
		err := next()

		if err == nil {
			return nil
		}

		// Json Parsing error
		var jsonErr *json.SyntaxError
		if errors.As(err, &jsonErr) {
			ctx.Response, _ = stackable.JsonResponse(
				http.StatusBadRequest,
				wrap.ErrorsResponse(err.Error(), err.Error()),
			)
			return nil
		}

		// Validation errors
		if errors.As(err, &validator.ValidationErrors{}) {
			ctx.Response, _ = stackable.JsonResponse(
				http.StatusUnprocessableEntity,
				wrap.ErrorsResponse(err.Error(), err.Error()),
			)
			return nil
		}

		// Api known errors
		var apiError *dto.ApiError
		if errors.As(err, &apiError) {
			ctx.Response, _ = stackable.JsonResponse(
				apiError.Status,
				wrap.ErrorsResponse(err.Error(), err.Error()),
			)
			return nil
		}

		// Generic error handler
		// Logging generic error
		ctx.Local.Logger.WithField("origin", "mw/errors_handler.go").WithField("err", err.Error()).Error("Unhandled generic error")

		ctx.Response, _ = stackable.JsonResponse(
			http.StatusInternalServerError,
			wrap.ErrorsResponse(err.Error(), err.Error()),
		)

		return nil
	},
)
