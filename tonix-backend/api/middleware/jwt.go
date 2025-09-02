package middleware

import (
	"net/http"
	"tonix/backend/api/context"
	"tonix/backend/api/dto"
	"tonix/backend/api/jwt"

	"github.com/saryginrodion/stackable"
)

var AccessJWTExtractor = stackable.WrapFunc(
	func(context *context.Context, next func() error) error {
		accessTokenString, err := context.Request.Cookie("access_token")

		if err != nil {
			return dto.NewApiError(http.StatusUnauthorized, "Failed to read access_token cookie.")
		}

		jwt.ParseAndVerifyToken[jwt.UserInfo](accessTokenString.Value, jwt.Access, context.Shared.Environment.JWT_SECRET)

		return next()
	},
)
