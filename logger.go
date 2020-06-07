package log

//Logger 日志接口
type Logger interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Error(v ...interface{})
	Warning(v ...interface{})
	Fatal(v ...interface{})

	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Warningf(format string, v ...interface{})
	Fatalf(format string, v ...interface{})

	SetLevel(level string)
}
