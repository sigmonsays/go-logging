// this generates much more efficient code
package logging

import "io"

type NullLogger struct{}

func NewNullLogger() Logger {
	return &NullLogger{}
}

func (l *NullLogger) SetWriter(io.Writer)                {}
func (l *NullLogger) Debug(args ...interface{})          {}
func (l *NullLogger) Info(args ...interface{})           {}
func (l *NullLogger) Warn(args ...interface{}) error     { return nil }
func (l *NullLogger) Error(args ...interface{}) error    { return nil }
func (l *NullLogger) Critical(args ...interface{}) error { return nil }
func (l *NullLogger) Trace(args ...interface{})          {}

func (l *NullLogger) Debugf(s string, args ...interface{})          {}
func (l *NullLogger) Infof(s string, args ...interface{})           {}
func (l *NullLogger) Warnf(s string, args ...interface{}) error     { return nil }
func (l *NullLogger) Errorf(s string, args ...interface{}) error    { return nil }
func (l *NullLogger) Criticalf(s string, args ...interface{}) error { return nil }
func (l *NullLogger) Tracef(s string, args ...interface{})          {}

// methods to implement the Logger interface
func (l *NullLogger) SetLevel(level string) error {
	return nil
}
func (l *NullLogger) GetLevel() string {
	return "disabled"
}

func (l *NullLogger) Close() {
}
func (l *NullLogger) Closed() bool {
	return false
}
func (l *NullLogger) Flush() {
}

func (l *NullLogger) IsTrace() bool    { return false }
func (l *NullLogger) IsDebug() bool    { return false }
func (l *NullLogger) IsInfo() bool     { return false }
func (l *NullLogger) IsWarn() bool     { return false }
func (l *NullLogger) IsError() bool    { return false }
func (l *NullLogger) IsCritical() bool { return false }
