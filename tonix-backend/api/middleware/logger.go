package middleware

import (
	"tonix/backend/api/context"
	"tonix/backend/logging"

	"github.com/radyshenkya/stackable"
	"github.com/sirupsen/logrus"
)

var LoggingMiddleware = stackable.WrapFunc(
	func(context *context.Context, next func() error) error {
		context.Local.Logger = logging.Logger("mw/logger.go").WithFields(
			logrus.Fields{
				"rid": context.Local.RequestId(),
				"ip":  context.Request.RemoteAddr,
			},
		)

		err := next()

		context.Local.Logger.Infof("%d - %s %s", context.Response.Status(), context.Request.Method, context.Request.URL.Path)

		return err
	},
)
