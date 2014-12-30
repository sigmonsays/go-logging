package logging
import "io"

type PrefixLogger struct {
	log    Logger
	Prefix string
}

func NewPrefixLogger(prefix string, l Logger) *PrefixLogger {
	logger := &PrefixLogger{
		log:    l,
		Prefix: prefix + " ",
	}
	return logger
}

func (p *PrefixLogger) SetWriter(io.Writer) {
}

func (p *PrefixLogger) SetLevel(level string) error {
	return nil
}
func (p *PrefixLogger) GetLevel() string {
	return p.log.GetLevel()
}

func (p *PrefixLogger) Tracef(format string, params ...interface{}) {
	p.log.Tracef(p.Prefix+format, params...)
}
func (p *PrefixLogger) Debugf(format string, params ...interface{}) {
	p.log.Debugf(p.Prefix+format, params...)
}
func (p *PrefixLogger) Infof(format string, params ...interface{}) {
	p.log.Infof(p.Prefix+format, params...)
}
func (p *PrefixLogger) Warnf(format string, params ...interface{}) error {
	return p.log.Warnf(p.Prefix+format, params...)
}
func (p *PrefixLogger) Errorf(format string, params ...interface{}) error {
	return p.log.Errorf(p.Prefix+format, params...)
}
func (p *PrefixLogger) Criticalf(format string, params ...interface{}) error {
	return p.log.Criticalf(p.Prefix+format, params...)
}

// this is a weird function but seems like the best way to insert
// a prefix argument in front of a array of arbitrary types...
func (p *PrefixLogger) prefix(v []interface{}) []interface{} {
	v2 := []interface{}{}
	v2 = append(v2, p.Prefix)
	v2 = append(v2, v...)
	return v2
}

func (p *PrefixLogger) Trace(v ...interface{}) {
	p.log.Trace(p.prefix(v)...)
}
func (p *PrefixLogger) Debug(v ...interface{}) {
	p.log.Debug(p.prefix(v)...)
}
func (p *PrefixLogger) Info(v ...interface{}) {
	p.log.Info(p.prefix(v)...)
}
func (p *PrefixLogger) Warn(v ...interface{}) error {
	return p.log.Warn(p.prefix(v)...)
}
func (p *PrefixLogger) Error(v ...interface{}) error {
	return p.log.Error(p.prefix(v)...)
}
func (p *PrefixLogger) Critical(v ...interface{}) error {
	return p.log.Critical(p.prefix(v)...)
}

func (p *PrefixLogger) Close() {
	p.log.Close()
}
func (p *PrefixLogger) Flush() {
	p.log.Flush()
}
func (p *PrefixLogger) Closed() bool {
	return p.log.Closed()
}

func (p *PrefixLogger) IsTrace() bool {
	return p.log.IsTrace()
}
func (p *PrefixLogger) IsDebug() bool {
	return p.log.IsDebug()
}
func (p *PrefixLogger) IsInfo() bool {
	return p.log.IsInfo()
}
func (p *PrefixLogger) IsWarn() bool {
	return p.log.IsWarn()
}
func (p *PrefixLogger) IsError() bool {
	return p.log.IsError()
}
func (p *PrefixLogger) IsCritical() bool {
	return p.log.IsCritical()
}


