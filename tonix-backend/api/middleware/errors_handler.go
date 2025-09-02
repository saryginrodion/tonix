package middleware

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"tonix/backend/api/context"
	"tonix/backend/api/dto"
	wrap "tonix/backend/api/dto/response_wrapper"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx"
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

		// PGX errors
		var pgxError pgx.PgError
		if errors.As(err, &pgxError) {
			switch pgxError.Code {
			// INVALID TEXT REPRESENTATION
			case "22P02":
				ctx.Response, _ = stackable.JsonResponse(
					http.StatusUnprocessableEntity,
					wrap.ErrorsResponse("Invalid Text Representation", err.Error()),
				)
				return nil
			}
		}

		// SQL no rows error
		if errors.Is(err, sql.ErrNoRows) {
			ctx.Response, _ = stackable.JsonResponse(
				http.StatusNotFound,
				wrap.ErrorsResponse("Not Found", err.Error()),
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
