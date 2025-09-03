package middleware

import (
	"net/http"
	"tonix/backend/api/context"
	"tonix/backend/api/dto"
	"tonix/backend/api/jwt"
	"tonix/backend/model"

	"github.com/saryginrodion/stackable"
)

var AccessJWTExtractor = stackable.WrapFunc(
	func(context *context.Context, next func() error) error {
		accessTokenString, err := context.Request.Cookie("access_token")

		if err != nil {
			return dto.NewApiError(http.StatusUnauthorized, "Failed to read access_token cookie.")
		}

		accessToken, err := jwt.ParseAndVerifyToken[jwt.UserInfo](accessTokenString.Value, jwt.Access, context.Shared.Environment.JWT_SECRET)

		if err != nil {
			return dto.NewApiError(http.StatusUnauthorized, "Failed to parse/verify access_token")
		}

		tokenWhitelist := model.TokenWhitelist(context.Shared.RedisClient)
		isExists, err := tokenWhitelist.IsWhitelisted(accessToken.Payload.TokenId)

		if err != nil {
			return err
		}

		if !isExists {
			return dto.NewApiError(http.StatusUnauthorized, "Token is not in whitelist")
		}

		context.Local.AccessJWT = accessToken

		return next()
	},
)

var RefreshJWTExtractor = stackable.WrapFunc(
	func(context *context.Context, next func() error) error {
		refreshTokenString, err := context.Request.Cookie("refresh_token")

		if err != nil {
			return dto.NewApiError(http.StatusUnauthorized, "Failed to read refresh_token cookie.")
		}

		refreshToken, err := jwt.ParseAndVerifyToken[jwt.UserInfo](refreshTokenString.Value, jwt.Refresh, context.Shared.Environment.JWT_SECRET)

		if err != nil {
			return dto.NewApiError(http.StatusUnauthorized, "Failed to parse/verify refresh_token")
		}

		tokenWhitelist := model.TokenWhitelist(context.Shared.RedisClient)
		isExists, err := tokenWhitelist.IsWhitelisted(refreshToken.Payload.TokenId)

		if err != nil {
			return err
		}

		if !isExists {
			return dto.NewApiError(http.StatusUnauthorized, "Token is not in whitelist")
		}

		context.Local.RefreshJWT = refreshToken

		return next()
	},
)
