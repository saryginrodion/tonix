package logging

import "github.com/sirupsen/logrus"

func baseLogger() *logrus.Logger {
	logger := logrus.New();

	return logger;
}

func Logger(origin string) *logrus.Entry {
	logger := baseLogger();

	return logger.WithField("origin", origin)
}
