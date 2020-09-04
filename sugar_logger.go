package helpers

type SugarLogger struct {
}

func (l SugarLogger) Warningf(message string, args ...interface{}) {
	zapLogger.Sugar().Warnf(message, args)
}

func (l SugarLogger) Errorf(message string, args ...interface{}) {
	zapLogger.Sugar().Errorf(message, args)
}

func (l SugarLogger) Debugf(message string, args ...interface{}) {
	zapLogger.Sugar().Debugf(message, args)
}

func (l SugarLogger) Infof(message string, args ...interface{}) {
	zapLogger.Sugar().Infof(message, args)
}

func GetLogger() SugarLogger {
	return SugarLogger{}
}
