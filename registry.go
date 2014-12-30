package logging

import (
	"fmt"
	"io"
	"os"
	"path"
	"sync"
)

type ReplaceFunction func(Logger)

var registry map[string]Logger

var replacefunction map[string]ReplaceFunction

var logOutput io.Writer
var outMux *sync.Mutex

func init() {
	registry = make(map[string]Logger)
	replacefunction = make(map[string]ReplaceFunction)
	logOutput = os.Stderr
	outMux = &sync.Mutex{}
}

func AddLogger(name string, log Logger, replacefunc ReplaceFunction) {
	Dbgf("AddLogger name=%s log=%#v replacefunc=%#v\n", name, log, replacefunc)

	if _, found := registry[name]; found {
		panic(fmt.Sprintf("AddLogger: Existing logger found", name))
	}
	outMux.Lock()
	log.SetWriter(logOutput)
	registry[name] = log
	replacefunction[name] = replacefunc
	outMux.Unlock()
}

func ReplaceLogger(name string, log Logger) {
	Dbgf("ReplaceLogger name=%s log=%#v\n", name, log)

	if GetLogger(name) == nil {
		AddLogger(name, log, nil)
	}

	if replace_function, ok := replacefunction[name]; ok {
		replace_function(log)
	} else {
		fmt.Println("ERROR REPLACING LOGGER:", name)
	}
}

func GetLogger(name string) Logger {
	if l, found := registry[name]; found {
		return l
	}
	return nil
}

type LogEntry struct {
	Name string
	Log  Logger
}

func ListLogger() map[string]Logger {
	return registry
}

var defaultLevel = "trace"

func DisableLog(name string) {
	ReplaceLogger(name, NewNullLogger())
}
func DisableLogs(pattern string) {
	for name, _ := range registry {
		if matched, _ := path.Match(pattern, name); matched {
			ReplaceLogger(name, NewNullLogger())
		}
	}
}
func DisableAllLogs() {
	for name, _ := range registry {
		println("** REPLACING ", name)
		ReplaceLogger(name, NewNullLogger())
	}
}

func SetLogLevel(level string) {
	defaultLevel = level
	for _, logger := range registry {
		logger.SetLevel(level)
	}
}
func SetLevel(name string, level string) {
	log := GetLogger(name)
	if log != nil {
		log.SetLevel(level)
	}
}
func SetLogLevels(levelMap map[string]string) {
	var log Logger
	for name, level := range levelMap {
		log = GetLogger(name)
		if log != nil {
			log.SetLevel(level)
		} else {
			println("*** Logger not found: ", name, " cannot set level ***")
		}
	}
}

func PrintLoggers() {
	fmt.Printf("Registered Loggers:\n")
	for name, logger := range registry {
		fmt.Printf("%s: %s\n", name, logger.GetLevel())
	}
}

func SetLogLevelsWithDefault(level string, levelMap map[string]string) {
	var to_set string
	var found bool
	for name, logger := range registry {
		to_set, found = levelMap[name]
		if !found {
			logger.SetLevel(level)
		} else {
			logger.SetLevel(to_set)
		}
	}
}

func SetLogOutput(out io.Writer) error {
	outMux.Lock()
	if out == nil {
		logOutput = os.Stderr
	} else {
		logOutput = out
	}
	outMux.Unlock()
	for _, logger := range registry {
		logger.SetWriter(logOutput)
	}
	return nil
}

// register a logger name
// if the name is not found, we use the standard logger
func Register(name string, replacefunc ReplaceFunction) (log Logger) {
	if replacefunc == nil {
		panic("You can't use nil replacefunc!")
	}
	log = GetLogger(name)
	if log == nil {
		// add a new standard logger at level trace until configured otherwise..
		log = NewStd2Logger3(defaultLevel, name)
		AddLogger(name, log, replacefunc)
	}
	return log
}
