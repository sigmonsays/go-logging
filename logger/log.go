package logger

import (
	"fmt"
)

type Logger interface {
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

	Message(Message)
}
type InnerLogger interface {
	FormatMessageLevel(lvl Level, format string, params ...interface{})
	ListMessageLevel(lvl Level, params ...interface{})
	Message(Message)
	//TODO: Flush() Idea: have a non-buffered chan for flush, that you check only when there are
	// no more messages in the messsage queue, then this could bubble down to Writer, and back up.
	//TODO: Update().Reset().WithConfig("foo.yaml").WithLevel().Update()
	// the last Update() is the builder create()/replace of the LogWriter.
	// its important to note that the Writer can only be swapped, once the existing filterchain is dried up.
	// so, this Update() method must
	// 1. spawn a goroutine to ReplaceWriter() on the Current FilterChain's tail, and close the oldWriter
	// 2. start the new one, and hook it up to the logger
	// 3. close() the existing filterchain
	// 4. once existing filterchain is done for, exit.
}
type Message interface {
	String() string
	Level() Level
}
type Formatter interface {
	CreateMessage(lvl Level, msg string, params ...interface{}) Message
	CreateMessageList(lvl Level, things ...interface{}) Message
}
type LogWriter interface {
	Write(Message)
}
type LogFilter interface {
	LogWriter
}

func Build() *Builder {
	return &Builder{formatter: &DefaultFormatter{}}
}

type Builder struct {
	lvl       Level
	output    LogWriter
	formatter Formatter
}

func (b *Builder) WithLevel(lvl Level) *Builder {
	b.lvl = lvl
	return b
}

func (b *Builder) WithOutput(output LogWriter) *Builder {
	b.output = output
	return b
}
func (b *Builder) WithFormatter(formatter Formatter) *Builder {
	b.formatter = formatter
	return b
}

func (b *Builder) Create() Logger {
	if b.output == nil {
		b.output = NewLogWriter()
	}
	var filter LogWriter
	if b.lvl != 0 {
		filter = NewLogFilter(b.lvl, b.output)
	} else {
		filter = b.output
	}
	return &LoggerWrapper{&LoggerImpl{filter, b.formatter}}
}

//TODO: define constants
type Level int

const (
	LvlCritical = Level(0)
	LvlError    = Level(1)
	LvlWarn     = Level(2)
	LvlInfo     = Level(3)
	LvlDebug    = Level(4)
	LvlTrace    = Level(5)
)

type LoggerWrapper struct {
	logger InnerLogger
}

func (l *LoggerWrapper) Message(m Message) {
	l.logger.Message(m)
}
func (l *LoggerWrapper) Criticalf(msg string, params ...interface{}) error {
	l.logger.FormatMessageLevel(LvlCritical, msg, params...)
	return fmt.Errorf(msg, params...)
}
func (l *LoggerWrapper) Errorf(msg string, params ...interface{}) error {
	l.logger.FormatMessageLevel(LvlError, msg, params...)
	return fmt.Errorf(msg, params...)
}
func (l *LoggerWrapper) Warnf(msg string, params ...interface{}) error {
	l.logger.FormatMessageLevel(LvlWarn, msg, params...)
	return fmt.Errorf(msg, params...)
}
func (l *LoggerWrapper) Infof(msg string, params ...interface{}) {
	l.logger.FormatMessageLevel(LvlInfo, msg, params...)
}
func (l *LoggerWrapper) Debugf(msg string, params ...interface{}) {
	l.logger.FormatMessageLevel(LvlDebug, msg, params...)
}
func (l *LoggerWrapper) Tracef(msg string, params ...interface{}) {
	l.logger.FormatMessageLevel(LvlTrace, msg, params...)
}

func (l *LoggerWrapper) Critical(params ...interface{}) error {
	l.logger.ListMessageLevel(LvlCritical, params...)
	return fmt.Errorf(fmt.Sprint(params...))
}
func (l *LoggerWrapper) Error(params ...interface{}) error {
	l.logger.ListMessageLevel(LvlError, params...)
	return fmt.Errorf(fmt.Sprint(params...))
}
func (l *LoggerWrapper) Warn(params ...interface{}) error {
	l.logger.ListMessageLevel(LvlWarn, params...)
	return fmt.Errorf(fmt.Sprint(params...))
}
func (l *LoggerWrapper) Info(params ...interface{}) {
	l.logger.ListMessageLevel(LvlInfo, params...)
}
func (l *LoggerWrapper) Debug(params ...interface{}) {
	l.logger.ListMessageLevel(LvlDebug, params...)
}
func (l *LoggerWrapper) Trace(params ...interface{}) {
	l.logger.ListMessageLevel(LvlTrace, params...)
}

type LoggerImpl struct {
	writer    LogWriter
	formatter Formatter
}

func (l *LoggerImpl) Message(m Message) {
	l.writer.Write(m)
}
func (l *LoggerImpl) FormatMessageLevel(lvl Level, msg string, params ...interface{}) {
	l.Message(l.formatter.CreateMessage(lvl, msg, params...))
}
func (l *LoggerImpl) ListMessageLevel(lvl Level, params ...interface{}) {
	l.Message(l.formatter.CreateMessageList(lvl, params...))
}

type FormatMessage struct {
	lvl    Level
	msg    string
	params []interface{}
}

func (f *FormatMessage) Level() Level {
	return f.lvl
}

func (f *FormatMessage) String() string {
	return fmt.Sprintf(f.msg, f.params...)
}

type ListMessage struct {
	lvl    Level
	params []interface{}
}

func (f *ListMessage) Level() Level {
	return f.lvl
}

func (f *ListMessage) String() string {
	return fmt.Sprint(f.params...)
}

//TODO: define more formatters to include lineNo, and funcname etc.
type DefaultFormatter struct {
}

func (f *DefaultFormatter) CreateMessage(lvl Level, msg string, params ...interface{}) Message {
	return &FormatMessage{lvl, msg, params}
}
func (f *DefaultFormatter) CreateMessageList(lvl Level, params ...interface{}) Message {
	return &ListMessage{lvl, params}
}
func Format(lvl Level, msg string, params ...interface{}) *FormatMessage {
	return &FormatMessage{lvl, msg, params}
}
