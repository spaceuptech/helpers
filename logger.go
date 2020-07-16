package helpers

import (
	"errors"
	"github.com/sirupsen/logrus"
)

const (
	logLevelDebug = "debug"
	logLevelInfo  = "info"
	logLevelError = "error"
)

var Logger = &logger{}

type logger struct {
}

func (l *logger) SetLogLevel(logLevel string) {
	switch logLevel {
	case logLevelDebug:
		logrus.SetLevel(logrus.DebugLevel)
	case logLevelInfo:
		logrus.SetLevel(logrus.InfoLevel)
	case logLevelError:
		logrus.SetLevel(logrus.ErrorLevel)
	default:
		logrus.Errorf("Invalid log level (%s) provided", logLevel)
		logrus.Infoln("Defaulting to `info` level")
		logrus.SetLevel(logrus.InfoLevel)
	}
}

// LogError logs the error in the proper format
func (l *logger) LogError(message, module, segment string, err error) error {
	// Prepare the fields
	entry := logrus.WithFields(logrus.Fields{"module": module, "segment": segment})
	if err != nil {
		entry = entry.WithError(err)
	}

	// Log the error
	entry.Errorln(message)

	// Return the error message
	return errors.New(message)
}

// LogWarn logs the warning message in the proper format
func (l *logger) LogWarn(message, module, segment string) {
	logrus.WithFields(logrus.Fields{"module": module, "segment": segment}).Warnln(message)
}

// LogInfo logs the info message in the proper format
func (l *logger) LogInfo(message, module, segment string) {
	logrus.WithFields(logrus.Fields{"module": module, "segment": segment}).Infoln(message)
}

// LogDebug logs the debug message in proper format
func (l *logger) LogDebug(message, module, segment string, extraFields map[string]interface{}) {
	entry := logrus.WithFields(logrus.Fields{"module": module, "segment": segment})
	if extraFields != nil {
		entry = entry.WithFields(extraFields)
	}
	entry.Debugln(message)
}
