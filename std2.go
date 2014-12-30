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

type Std2Logger struct {
	CallDepth int
	mu        sync.RWMutex
	level     int
	name      string
	trace     *log.Logger
	debug     *log.Logger
	info      *log.Logger
	warn      *log.Logger
	err       *log.Logger
	critical  *log.Logger
}

// standard logger
func NewStd2Logger(level string) *Std2Logger {
	return NewStd2Logger3(level, "")
}

func NewStd2Logger3(level string, name string) *Std2Logger {
	lvl, err := LevelFromString(level)
	if err != nil {
		panic(fmt.Sprintf("invalid level: %s", level))
	}
	l := &Std2Logger{
		CallDepth: 2,
		level:     lvl,
		name:      name,
	}
	l.SetWriter(os.Stderr)
	return l
}

func (l *Std2Logger) SetWriter(out io.Writer) {
	l.trace = log.New(out, l.name+" TRACE ", log.Lshortfile)
	l.debug = log.New(out, l.name+" DEBUG ", log.Lshortfile)
	l.info = log.New(out, l.name+" INFO ", log.Lshortfile)
	l.warn = log.New(out, l.name+" WARN ", log.Lshortfile)
	l.err = log.New(out, l.name+" ERROR ", log.Lshortfile)
	l.critical = log.New(out, l.name+" CRITICAL ", log.Lshortfile)
}

func (l *Std2Logger) GetLevel() string {
	l.mu.Lock()
	defer l.mu.Unlock()
	return Levels[l.level]
}
func (l *Std2Logger) setlevel(level int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = level
}

// print style functions
func (l *Std2Logger) Trace(args ...interface{}) {
	if l.level <= TRACE {
		l.trace.Output(l.CallDepth, fmt.Sprintln(args...))
	}
}
func (l *Std2Logger) Debug(args ...interface{}) {
	if l.level <= DEBUG {
		l.debug.Output(l.CallDepth, fmt.Sprintln(args...))
	}
}
func (l *Std2Logger) Info(args ...interface{}) {
	if l.level <= INFO {
		l.info.Output(l.CallDepth, fmt.Sprintln(args...))
	}
}
func (l *Std2Logger) Warn(args ...interface{}) error {
	if l.level <= WARNING {
		msg := fmt.Sprintln(args...)
		l.warn.Output(l.CallDepth, msg)
		return errors.New(msg)
	}
	return errors.New(fmt.Sprintln(args...))
}
func (l *Std2Logger) Error(args ...interface{}) error {
	if l.level <= ERROR {
		msg := fmt.Sprintln(args...)
		l.err.Output(l.CallDepth, msg)
		return errors.New(msg)
	}
	return errors.New(fmt.Sprintln(args...))
}
func (l *Std2Logger) Critical(args ...interface{}) error {
	if l.level <= CRITICAL {
		msg := fmt.Sprintln(args...)
		l.critical.Output(l.CallDepth, msg)
		return errors.New(msg)
	}
	return errors.New(fmt.Sprintln(args...))
}

// printf style functions
func (l *Std2Logger) Tracef(format string, args ...interface{}) {
	if l.level <= TRACE {
		l.trace.Output(l.CallDepth, fmt.Sprintf(format, args...))
	}
}
func (l *Std2Logger) Debugf(format string, args ...interface{}) {
	if l.level <= DEBUG {
		l.debug.Output(l.CallDepth, fmt.Sprintf(format, args...))
	}
}
func (l *Std2Logger) Infof(format string, args ...interface{}) {
	if l.level <= INFO {
		l.info.Output(l.CallDepth, fmt.Sprintf(format, args...))
	}
}
func (l *Std2Logger) Warnf(format string, args ...interface{}) error {
	if l.level <= WARNING {
		msg := fmt.Sprintf(format, args...)
		l.warn.Output(l.CallDepth, msg)
		return errors.New(msg)
	}
	return fmt.Errorf(format, args...)
}
func (l *Std2Logger) Errorf(format string, args ...interface{}) error {
	if l.level <= ERROR {
		msg := fmt.Sprintf(format, args...)
		l.err.Output(l.CallDepth, msg)
		return errors.New(msg)
	}
	return fmt.Errorf(format, args...)
}
func (l *Std2Logger) Criticalf(format string, args ...interface{}) error {
	if l.level <= CRITICAL {
		msg := fmt.Sprintf(format, args...)
		l.critical.Output(l.CallDepth, msg)
		return errors.New(msg)
	}
	return fmt.Errorf(format, args...)
}

// methods to implement the Logger interface
func (l *Std2Logger) SetLevel(level string) error {
	lvl, err := LevelFromString(level)
	if err != nil {
		return fmt.Errorf("SetLevel: %s", err)
	}
	l.setlevel(lvl)
	return nil
}

func (l *Std2Logger) Close() {
}
func (l *Std2Logger) Closed() bool {
	return false
}
func (l *Std2Logger) Flush() {
}

func (l *Std2Logger) IsTrace() bool {
	return l.level <= TRACE
}

func (l *Std2Logger) IsDebug() bool {
	return l.level <= DEBUG
}

func (l *Std2Logger) IsInfo() bool {
	return l.level <= INFO
}

func (l *Std2Logger) IsWarn() bool {
	return l.level <= WARNING
}

func (l *Std2Logger) IsError() bool {
	return l.level <= ERROR
}

func (l *Std2Logger) IsCritical() bool {
	return l.level <= CRITICAL
}
