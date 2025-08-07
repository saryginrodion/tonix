package context

import (
	"tonix/backend/logging"

	"github.com/saryginrodion/stackable/middleware"
	"github.com/sirupsen/logrus"
)

type LocalState struct {
	*middleware.LocalRequestId
	Logger *logrus.Entry
}

func (l LocalState) Default() any {
	return LocalState{
		LocalRequestId: &middleware.LocalRequestId{},
		Logger: logging.LoggerWithOrigin("local.go"),
	}
}
