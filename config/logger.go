package config

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func initColors() {
	if runtime.GOOS == "windows" {
		Reset = ""
		Red = ""
		Green = ""
		Yellow = ""
		Blue = ""
		Purple = ""
		Cyan = ""
		Gray = ""
		White = ""
	}
}

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger(prefix string) *Logger {
	initColors()
	writer := io.Writer(os.Stdout)

	logger := log.New(writer, prefix, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(logger.Writer(), fmt.Sprintf(Purple+"DEBUG: On %v → "+Reset, logger.Prefix()), logger.Flags()),
		info:    log.New(logger.Writer(), fmt.Sprintf(Cyan+"INFO: On %v → "+Reset, logger.Prefix()), logger.Flags()),
		warning: log.New(logger.Writer(), fmt.Sprintf(Yellow+"WARNING: On %v → "+Reset, logger.Prefix()), logger.Flags()),
		err:     log.New(logger.Writer(), fmt.Sprintf(Red+"ERROR: On %v → "+Reset, logger.Prefix()), logger.Flags()),
		writer:  logger.Writer(),
	}
}

func (l Logger) Debug(values ...interface{}) {
	l.debug.Println(values...)
}

func (l Logger) Info(values ...interface{}) {
	l.info.Println(values...)
}

func (l Logger) Warn(values ...interface{}) {
	l.warning.Println(values...)
}

func (l Logger) Error(values ...interface{}) {
	l.err.Println(values...)
}

func (l Logger) Debugf(format string, values ...interface{}) {
	l.debug.Printf(format, values...)
}

func (l Logger) Infof(format string, values ...interface{}) {
	l.info.Printf(format, values...)
}

func (l Logger) Warnf(format string, values ...interface{}) {
	l.warning.Printf(format, values...)
}

func (l Logger) Errorf(format string, values ...interface{}) {
	l.err.Printf(format, values...)
}
