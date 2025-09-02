package v1

import (
	"net/http"
	"time"
	"tonix/backend/api/context"
	"tonix/backend/api/dto"
	"tonix/backend/api/dto/requests"
	wrap "tonix/backend/api/dto/response_wrapper"
	"tonix/backend/api/dto/view"
	"tonix/backend/api/jwt"
	"tonix/backend/api/utils"
	"tonix/backend/env_vars"
	"tonix/backend/model"

	"github.com/gofrs/uuid"
	"github.com/saryginrodion/stackable"
)

type TokenPair struct {
	Access  jwt.TokenPayload[jwt.UserInfo]
	Refresh jwt.TokenPayload[jwt.UserInfo]
}

func CreateTokenPair(user *model.User, envVars env_vars.EnvVars) (TokenPair, error) {
	tokenUuid, _ := uuid.NewV4()
	userInfo := jwt.UserInfo{
		Uid: user.Id,
	}

	accessExpTime := time.Now()
	accessExpTime.Add(envVars.JWT_ACCESS_COOLDOWN_DURATION)
	accessTokenPayload := jwt.NewTokenPayload(userInfo, accessExpTime, jwt.Access, tokenUuid.String())

	refreshExpTime := time.Now()
	refreshExpTime.Add(envVars.JWT_REFRESH_COOLDOWN_DURATION)
	refreshTokenPayload := jwt.NewTokenPayload(userInfo, refreshExpTime, jwt.Refresh, tokenUuid.String())

	return TokenPair{
		Access:  accessTokenPayload,
		Refresh: refreshTokenPayload,
	}, nil
}

var Registration = stackable.WrapFunc(
	func(ctx *context.Context, next func() error) error {
		body, err := utils.ParseAndValidateJson(ctx.Request.Body, requests.RegistrationBody{})
		if err != nil {
			return err
		}

		users := model.Users(ctx.Shared.DB)

		isRegistered, err := users.IsRegistered(body.Username, body.Email)
		if err != nil {
			return err
		}

		if isRegistered {
			return dto.NewApiError(http.StatusConflict, "User with this email or username is already registered")
		}

		hashedPassword, err := utils.HashPassword(body.Password)
		if err != nil {
			return err
		}

		id, err := users.New(body.Email, *hashedPassword, body.Username)
		if err != nil {
			return err
		}

		user, err := users.ById(*id)
		if err != nil {
			return err
		}

		// Creating cookie pair
		tokenPair, err := CreateTokenPair(user, ctx.Shared.Environment)
		if err != nil {
			return err
		}

		accessToken, err := jwt.GenerateToken(tokenPair.Access, ctx.Shared.Environment.JWT_SECRET)
		refreshToken, err := jwt.GenerateToken(tokenPair.Refresh, ctx.Shared.Environment.JWT_SECRET)

		if err != nil {
			return err
		}

		accessCookie := http.Cookie{
			Name:     "access_token",
			Value:    accessToken,
			Path:     "/",
			Secure:   false,
			HttpOnly: true,
			SameSite: 0,
		}

		refreshCookie := http.Cookie{
			Name:     "refresh_token",
			Value:    refreshToken,
			Path:     "/",
			Secure:   false,
			HttpOnly: true,
			SameSite: 0,
		}

		resp, _ := stackable.JsonResponse(
			http.StatusOK,
			wrap.OkResponse(view.UserToView(user)),
		)

		headers := resp.Headers()
		headers.Add("Set-Cookie", accessCookie.String())
		headers.Add("Set-Cookie", refreshCookie.String())

		resp.SetHeaders(headers)

		ctx.Response = resp

		return next()
	},
)

var Login = stackable.WrapFunc(
	func(ctx *context.Context, next func() error) error {
		body, err := utils.ParseAndValidateJson(ctx.Request.Body, requests.LoginBody{})
		if err != nil {
			return err
		}

		users := model.Users(ctx.Shared.DB)
		user, err := users.ByEmail(body.Email)
		if err != nil {
			return err
		}

		if !utils.ComparePassword(body.Password, user.Password) {
			return dto.NewApiError(http.StatusUnauthorized, "Invalid email or password")
		}

		// Creating cookie pair
		tokenPair, err := CreateTokenPair(user, ctx.Shared.Environment)
		if err != nil {
			return err
		}

		accessToken, err := jwt.GenerateToken(tokenPair.Access, ctx.Shared.Environment.JWT_SECRET)
		refreshToken, err := jwt.GenerateToken(tokenPair.Refresh, ctx.Shared.Environment.JWT_SECRET)

		if err != nil {
			return err
		}

		accessCookie := http.Cookie{
			Name:     "access_token",
			Value:    accessToken,
			Path:     "/",
			Secure:   false,
			HttpOnly: true,
			SameSite: 0,
		}

		refreshCookie := http.Cookie{
			Name:     "refresh_token",
			Value:    refreshToken,
			Path:     "/",
			Secure:   false,
			HttpOnly: true,
			SameSite: 0,
		}

		resp, _ := stackable.JsonResponse(
			http.StatusOK,
			wrap.OkResponse(view.UserToView(user)),
		)

		headers := resp.Headers()
		headers.Add("Set-Cookie", accessCookie.String())
		headers.Add("Set-Cookie", refreshCookie.String())

		resp.SetHeaders(headers)

		ctx.Response = resp

		return next()
	},
)
