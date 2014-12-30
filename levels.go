package logging

import (
	"fmt"
	"strings"
)

// TRACE 0, DEBUG 10, INFO 20, WARNING 30, ERROR 40, CRITICAL 50
const (
	TRACE = iota * 10
	DEBUG
	INFO
	WARNING
	ERROR
	CRITICAL
)

var Levels = map[int]string{
	TRACE:    "TRACE",
	DEBUG:    "DEBUG",
	INFO:     "INFO",
	WARNING:  "WARNING",
	ERROR:    "ERROR",
	CRITICAL: "CRITICAL",
}

var AliasLevels = map[string]int{
	"WARN": WARNING,
}

var Constants = make(map[string]int, len(Levels))
var Std Logger

func init() {
	for k, v := range Levels {
		Constants[v] = k
	}
	for k, v := range AliasLevels {
		Constants[k] = v
	}
	Std = NewStandardLogger("WARNING")
}

func LevelFromString(level string) (int, error) {
	ulevel := strings.ToUpper(level)
	v, ok := Constants[ulevel]
	if !ok {
		fmt.Printf("Invalid Log Level: %s, valid levels %#v\n", ulevel, Constants)
		return 0, fmt.Errorf("Invalid level: %s, valid levels %#v", ulevel, Constants)
	}
	return v, nil
}
