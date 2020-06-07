package log

import (
	"github.com/sirupsen/logrus"
)

type logrusLog struct {
	logger *logrus.Logger
}

func newTextLog() *logrusLog {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	return &logrusLog{
		logger: logger,
	}
}

func (l *logrusLog) Debug(v ...interface{}) {
	l.logger.Debug(v...)
}

func (l *logrusLog) Info(v ...interface{}) {
	l.logger.Info(v...)
}

func (l *logrusLog) Warning(v ...interface{}) {
	l.logger.Warning(v...)
}

func (l *logrusLog) Error(v ...interface{}) {
	l.logger.Error(v...)
}

func (l *logrusLog) Fatal(v ...interface{}) {
	l.logger.Fatal(v...)
}

func (l *logrusLog) Debugf(format string, v ...interface{}) {
	l.logger.Debugf(format, v...)
}

func (l *logrusLog) Infof(format string, v ...interface{}) {
	l.logger.Infof(format, v...)
}

func (l *logrusLog) Warningf(format string, v ...interface{}) {
	l.logger.Warningf(format, v...)
}

func (l *logrusLog) Errorf(format string, v ...interface{}) {
	l.logger.Errorf(format, v...)
}

func (l *logrusLog) Fatalf(format string, v ...interface{}) {
	l.logger.Fatalf(format, v...)
}

func (l *logrusLog) SetLevel(level string) {
	var lv logrus.Level
	switch level {
	case "debug":
		lv = logrus.DebugLevel
	case "info":
		lv = logrus.InfoLevel
	case "warn":
		lv = logrus.WarnLevel
	case "error":
		lv = logrus.ErrorLevel
	case "fatal":
		lv = logrus.FatalLevel
	default:
		panic("unknow log level:" + level)
	}
	l.logger.SetLevel(lv)
}

var defaultLog = newTextLog()

//Debug Debug级别日志
func Debug(v ...interface{}) {
	defaultLog.Debug(v...)
}

//Info Info级别日志
func Info(v ...interface{}) {
	defaultLog.Info(v...)
}

//Warning Warning级别日志
func Warning(v ...interface{}) {
	defaultLog.Warning(v...)
}

//Error Error级别日志
func Error(v ...interface{}) {
	defaultLog.Error(v...)
}

//Fatal Fatal级别日志
func Fatal(v ...interface{}) {
	defaultLog.Fatal(v...)
}

//Debugf Debug级别日志
func Debugf(format string, v ...interface{}) {
	defaultLog.Debugf(format, v...)
}

//Infof Info级别日志
func Infof(format string, v ...interface{}) {
	defaultLog.Infof(format, v...)
}

//Warningf Warning级别日志
func Warningf(format string, v ...interface{}) {
	defaultLog.Warningf(format, v...)
}

//Errorf Error级别日志
func Errorf(format string, v ...interface{}) {
	defaultLog.Errorf(format, v...)
}

//Fatalf Fatal级别日志
func Fatalf(format string, v ...interface{}) {
	defaultLog.Fatalf(format, v...)
}

//SetLevel 设置日志输出级别
func SetLevel(level string) {
	defaultLog.SetLevel(level)
}

//GetLogger 返回自身
func GetLogger() Logger {
	return defaultLog
}

//New 创建logger实例
func New() Logger {
	return newTextLog()
}
