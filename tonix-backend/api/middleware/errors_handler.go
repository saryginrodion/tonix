package middleware

import (
	"net/http"
	"tonix/backend/api/context"
	wrap "tonix/backend/api/dto/response_wrapper"
	"tonix/backend/logging"

	"github.com/radyshenkya/stackable"
)

var log = logging.LoggerWithOrigin("mw/errors_handler.go")

var ErrorsHandlerMiddleware = stackable.WrapFunc(
	func(ctx *context.Context, next func() error) error {
		err := next()

		if err == nil {
			return nil
		}

		// Generic error handler
		// Logging generic error
		ctx.Local.Logger.WithField("origin", "mw/errors_handler.go").WithField("err", err.Error()).Error("Unhandled generic error")

		ctx.Response, _ = stackable.JsonResponse(
			http.StatusInternalServerError,
			wrap.ErrorsResponse(err.Error(), []string{err.Error()}),
		)

		return nil
	},
)
