package logging

import (
	"fmt"
)

// this is here to help in tracing down what the logging libary does

var (
	DebugLoggingLibrary = false
)

func Dbgf(s string, args ...interface{}) {
	if DebugLoggingLibrary {
		fmt.Printf("[DEBUG go-logging] "+s, args...)
	}
}
