// standard logger interface which inherits the normal log package to provide log levels
//
package logging

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

type StandardLogger struct {
	CallDepth int
	mu        sync.RWMutex
	level     int
	*log.Logger
}

// standard logger
func NewStandardLogger(level string) *StandardLogger {
	lvl, err := LevelFromString(level)
	if err != nil {
		panic(fmt.Sprintf("invalid level: %s", level))
	}
	return NewStandardLogger2(lvl)
}

func NewStandardLogger2(level int) *StandardLogger {
	l := &StandardLogger{
		Logger:    log.New(os.Stderr, "", log.Lshortfile),
		CallDepth: 4,
		level:     level,
	}

	return l
}

func NewStandardLogger3(level, depth int) *StandardLogger {
	l := &StandardLogger{
		Logger:    log.New(os.Stderr, "", log.Lshortfile),
		CallDepth: depth,
		level:     level,
	}

	return l
}

func (l *StandardLogger) SetWriter(out io.Writer) {
	l.Logger = log.New(out, "", log.Lshortfile)
}

func (l *StandardLogger) GetLevel() string {
	l.mu.Lock()
	defer l.mu.Unlock()
	return Levels[l.level]
}
func (l *StandardLogger) setlevel(level int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = level
}
func (l *StandardLogger) LogLine(level int, args ...interface{}) error {
	l.mu.RLock()
	defer l.mu.RUnlock()
	if l.level <= level {
		return l.Output(l.CallDepth, Levels[level]+" "+fmt.Sprintln(args...))
	}
	return nil
}
func (l *StandardLogger) Output(calldepth int, s string) error {
	return l.Logger.Output(calldepth, s)
}

// print style functions
func (l *StandardLogger) Trace(args ...interface{}) {
	l.LogLine(TRACE, args...)
}
func (l *StandardLogger) Debug(args ...interface{}) {
	l.LogLine(DEBUG, args...)
}
func (l *StandardLogger) Info(args ...interface{}) {
	l.LogLine(INFO, args...)
}
func (l *StandardLogger) Warn(args ...interface{}) error {
	l.LogLine(WARNING, args...)
	return errors.New(fmt.Sprintln(args...))
}
func (l *StandardLogger) Error(args ...interface{}) error {
	l.LogLine(ERROR, args...)
	return errors.New(fmt.Sprintln(args...))
}
func (l *StandardLogger) Critical(args ...interface{}) error {
	l.LogLine(CRITICAL, args...)
	return errors.New(fmt.Sprintln(args...))
}

// printf style functions
func (l *StandardLogger) Tracef(format string, args ...interface{}) {
	l.LogLine(TRACE, fmt.Sprintf(format, args...))
}
func (l *StandardLogger) Debugf(format string, args ...interface{}) {
	l.LogLine(DEBUG, fmt.Sprintf(format, args...))
}
func (l *StandardLogger) Infof(format string, args ...interface{}) {
	l.LogLine(INFO, fmt.Sprintf(format, args...))
}
func (l *StandardLogger) Warnf(format string, args ...interface{}) error {
	l.LogLine(WARNING, fmt.Sprintf(format, args...))
	return fmt.Errorf(format, args...)
}
func (l *StandardLogger) Errorf(format string, args ...interface{}) error {
	l.LogLine(ERROR, fmt.Sprintf(format, args...))
	return fmt.Errorf(format, args...)
}
func (l *StandardLogger) Criticalf(format string, args ...interface{}) error {
	l.LogLine(CRITICAL, fmt.Sprintf(format, args...))
	return fmt.Errorf(format, args...)
}

// helper functions to use the provided "standard" logger
// print style functions helper functions
func Trace(args ...interface{}) {
	Std.Trace(args...)
}
func Debug(args ...interface{}) {
	Std.Debug(args...)
}
func Info(args ...interface{}) {
	Std.Info(args...)
}
func Warn(args ...interface{}) error {
	Std.Warn(args...)
	return errors.New(fmt.Sprintln(args...))
}
func Error(args ...interface{}) error {
	Std.Error(args...)
	return errors.New(fmt.Sprintln(args...))
}
func Critical(args ...interface{}) error {
	Std.Critical(args...)
	return errors.New(fmt.Sprintln(args...))
}

// printf style functions helper functions
func Tracef(format string, args ...interface{}) {
	Std.Tracef(format, args...)
}
func Debugf(format string, args ...interface{}) {
	Std.Debugf(format, args...)
}
func Infof(format string, args ...interface{}) {
	Std.Infof(format, args...)
}
func Warnf(format string, args ...interface{}) error {
	Std.Warnf(format, args...)
	return fmt.Errorf(format, args...)
}
func Errorf(format string, args ...interface{}) error {
	Std.Errorf(format, args...)
	return fmt.Errorf(format, args...)
}
func Criticalf(format string, args ...interface{}) error {
	return Std.Criticalf(format, args...)
	return fmt.Errorf(format, args...)
}

// methods to implement the Logger interface
func (l *StandardLogger) SetLevel(level string) error {
	lvl, err := LevelFromString(level)
	if err != nil {
		return fmt.Errorf("SetLevel: %s", err)
	}
	l.setlevel(lvl)
	return nil
}

func (l *StandardLogger) Close() {
}
func (l *StandardLogger) Closed() bool {
	return false
}
func (l *StandardLogger) Flush() {
}

func (l *StandardLogger) IsTrace() bool {
	return l.level <= TRACE
}

func (l *StandardLogger) IsDebug() bool {
	return l.level <= DEBUG
}

func (l *StandardLogger) IsInfo() bool {
	return l.level <= INFO
}

func (l *StandardLogger) IsWarn() bool {
	return l.level <= WARNING
}

func (l *StandardLogger) IsError() bool {
	return l.level <= ERROR
}

func (l *StandardLogger) IsCritical() bool {
	return l.level <= CRITICAL
}
