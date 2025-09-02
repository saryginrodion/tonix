package context

import (
	"tonix/backend/api/jwt"
	"tonix/backend/logging"

	"github.com/saryginrodion/stackable/middleware"
	"github.com/sirupsen/logrus"
)

type LocalState struct {
	*middleware.LocalRequestId
	Logger *logrus.Entry
	AccessJWT *jwt.Token[jwt.UserInfo]
}

func (l LocalState) Default() any {
	return LocalState{
		LocalRequestId: &middleware.LocalRequestId{},
		Logger: logging.LoggerWithOrigin("local.go"),
		AccessJWT: nil,
	}
}
