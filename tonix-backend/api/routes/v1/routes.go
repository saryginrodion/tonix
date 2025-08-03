package v1

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"tonix/backend/api/context"
	"tonix/backend/api/dto"

	"github.com/radyshenkya/stackable"
)

var GetIndex = stackable.WrapFunc(
	func(ctx *context.Context, next func() error) error {
		if ctx.Request.URL.Path != "/api/v1/" || ctx.Request.Method != "GET" {
			ctx.Response = stackable.NewHttpResponse(
				http.StatusNotFound,
				"text/html",
				"<h1>404 - Not found</h1>",
			)

			return next()
		}

		ctx.Response = stackable.NewHttpResponse(
			http.StatusOK,
			"text/html",
			fmt.Sprintf("<h1>Hello world!</h1>POSTGRES_CONNECTION_URL=%s", ctx.Shared.Environment.POSTGRES_CONNECTION_URL),
		)

		return next()
	},
)

var PostTestMessage = stackable.WrapFunc(
	func(ctx *context.Context, next func() error) error {
		body, _ := io.ReadAll(ctx.Request.Body)
		reqMsg := dto.TestMessage{}

		json.Unmarshal(body, &reqMsg)

		ctx.Response, _ = stackable.JsonResponse(
			http.StatusOK,
			dto.TestMessage{Message: reqMsg.Message},
		)

		return next()
	},
)
