package logger

func NewLogFilter(lvl Level, writer LogWriter) LogWriter {
	ch := make(chan Message, 200)
	return &LogFilterLevel{ch, lvl, writer, false}
}

type LogFilterLevel struct {
	input   chan Message
	lvl     Level
	writer  LogWriter
	started bool
}

func (l *LogFilterLevel) Write(m Message) {
	if !l.started {
		go l.Start()
	}
	l.input <- m
}
func (l *LogFilterLevel) Start() {
	for {
		select {
		case message, _ := <-l.input:
			if message.Level() >= l.lvl {
				l.writer.Write(message)
			}
		}
	}
}

func NewLogWriter() LogWriter {
	ch := make(chan Message, 200)
	return &LogWriterStdout{ch: ch, started: false}
}

type LogWriterStdout struct {
	ch      chan Message
	started bool
}

func (c *LogWriterStdout) Write(m Message) {
	if !c.started {
		go c.Start()
	}
	c.ch <- m
}

func (c *LogWriterStdout) MessageChan() chan Message {
	return c.ch
}

func (c *LogWriterStdout) Start() {
	for {
		select {
		case message, _ := <-c.ch:
			println(message.String())
		}
	}
}
