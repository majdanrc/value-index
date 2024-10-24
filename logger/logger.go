package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

var logPrefixes = map[LogLevel]string{
	DEBUG: "[DEBUG] ",
	INFO:  "[INFO] ",
	WARN:  "[WARN] ",
	ERROR: "[ERROR] ",
}

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}

type StdLogger struct {
	level  LogLevel
	logger *log.Logger
}

func NewStdLogger(level LogLevel, output io.Writer) *StdLogger {
	if output == nil {
		output = os.Stdout
	}
	return &StdLogger{
		level:  level,
		logger: log.New(output, "", log.LstdFlags),
	}
}

func (l *StdLogger) formatAndLog(level LogLevel, msg string) {
	if l.level <= level {
		l.logger.Println(logPrefixes[level] + msg)
	}
}

func (l *StdLogger) Debug(msg string) {
	l.formatAndLog(DEBUG, msg)
}

func (l *StdLogger) Info(msg string) {
	l.formatAndLog(INFO, msg)
}

func (l *StdLogger) Warn(msg string) {
	l.formatAndLog(WARN, msg)
}

func (l *StdLogger) Error(msg string) {
	l.formatAndLog(ERROR, msg)
}

func LogLevelFromString(level string) (LogLevel, error) {
	switch strings.ToLower(level) {
	case "debug":
		return DEBUG, nil
	case "info":
		return INFO, nil
	case "warn":
		return WARN, nil
	case "error":
		return ERROR, nil
	default:
		return INFO, fmt.Errorf("invalid log level: %s", level)
	}
}
