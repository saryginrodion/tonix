package view

import (
	"time"
	"tonix/backend/api/utils"
	"tonix/backend/model"
)

type SelfUserView struct {
	Id                 string    `json:"id"`
	Email              string    `json:"email"`
	Username           string    `json:"username"`
	DisplayableName    string    `json:"displayable_name"`
	Description        string    `json:"description"`
	EmailVerified      bool      `json:"email_verified"`
	Balance            int32     `json:"balance"`
	WithdrawalBalance  int32     `json:"withdrawal_balance"`
	LastWithdrawalCard string    `json:"last_withdrawal_card"`
	AvatarId           *string   `json:"avatar_id"`
	IdentityPhotoId    *string   `json:"identity_photo_id"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

func ToSelfUserView(u *model.User) SelfUserView {
	return SelfUserView{
		Id:                 u.Id,
		Email:              u.Email,
		Username:           u.Username,
		DisplayableName:    u.DisplayableName,
		Description:        u.Description,
		EmailVerified:      u.EmailVerified,
		Balance:            u.Balance,
		WithdrawalBalance:  u.WithdrawalBalance,
		LastWithdrawalCard: u.LastWithdrawalCard,
		AvatarId:           utils.NullableToString(u.AvatarId),
		IdentityPhotoId:    utils.NullableToString(u.IdentityPhotoId),
		CreatedAt:          u.CreatedAt,
		UpdatedAt:          u.UpdatedAt,
	}
}

type UserView struct {
	Id              string    `json:"id"`
	Username        string    `json:"username"`
	DisplayableName string    `json:"displayable_name"`
	Description     string    `json:"description"`
	AvatarId        *string   `json:"avatar_id"`
	CreatedAt       time.Time `json:"created_at"`
}

func ToUserView(u *model.User) UserView {
	return UserView{
		Id:              u.Id,
		Username:        u.Username,
		DisplayableName: u.DisplayableName,
		Description:     u.Description,
		CreatedAt:       u.CreatedAt,
		AvatarId:        utils.NullableToString(u.AvatarId),
	}
}
