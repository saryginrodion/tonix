package v1

import (
	"net/http"
	"tonix/backend/api/context"
	"tonix/backend/api/dto/view"
	"tonix/backend/model"

	"github.com/saryginrodion/stackable"
)

var ProfileSelf = stackable.WrapFunc(
	func(ctx *context.Context, next func() error) error {
		users := model.Users(ctx.Shared.DB)
		userId := ctx.Local.AccessJWT.Payload.Data.Uid

		user, err := users.ById(model.Id(userId))
		if err != nil {
			return err
		}

		ctx.Response, _ = stackable.JsonResponse(http.StatusOK, view.ToSelfUserView(user))

		return next()
	},
)

var Profile = stackable.WrapFunc(
	func(ctx *context.Context, next func() error) error {
		users := model.Users(ctx.Shared.DB)
		userId := ctx.Request.PathValue("id")

		user, err := users.ById(model.Id(userId))
		if err != nil {
			return err
		}

		ctx.Response, _ = stackable.JsonResponse(http.StatusOK, view.ToUserView(user))

		return next()
	},
)
