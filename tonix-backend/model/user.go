package model

import (
	"database/sql"
	"time"
)

type User struct {
	Id                 string         `db:"id"`
	Email              string         `db:"email"`
	Password           string         `db:"password"`
	Username           string         `db:"username"`
	DisplayableName    string         `db:"displayable_name"`
	Description        string         `db:"description"`
	EmailVerified      bool           `db:"email_verified"`
	Balance            int32          `db:"balance"`
	WithdrawalBalance  int32          `db:"withdrawal_balance"`
	LastWithdrawalCard string         `db:"last_withdrawal_card"`
	AvatarId           sql.NullString `db:"avatar_id"`
	IdentityPhotoId    sql.NullString `db:"identity_photo_id"`
	UpdatedAt          time.Time      `db:"updated_at"`
	CreatedAt          time.Time      `db:"created_at"`
}
