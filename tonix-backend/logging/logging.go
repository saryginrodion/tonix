package logging

import "github.com/sirupsen/logrus"

func baseLogger() *logrus.Logger {
	logger := logrus.New();

	return logger;
}

func LoggerWithOrigin(origin string) *logrus.Entry {
	logger := baseLogger();

	return logger.WithField("origin", origin)
}

func Logger() *logrus.Logger {
	return baseLogger()
}
