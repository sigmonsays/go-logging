package logging

import "io"

type Logger interface {
	GetLevel() string
	SetLevel(level string) error
	SetWriter(io.Writer)

	Tracef(format string, params ...interface{})
	Debugf(format string, params ...interface{})
	Infof(format string, params ...interface{})
	Warnf(format string, params ...interface{}) error
	Errorf(format string, params ...interface{}) error
	Criticalf(format string, params ...interface{}) error

	Trace(v ...interface{})
	Debug(v ...interface{})
	Info(v ...interface{})
	Warn(v ...interface{}) error
	Error(v ...interface{}) error
	Critical(v ...interface{}) error

	Close()
	Flush()
	Closed() bool

	IsTrace() bool
	IsDebug() bool
	IsInfo() bool
	IsWarn() bool
	IsError() bool
	IsCritical() bool
}
