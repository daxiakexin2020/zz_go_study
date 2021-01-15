package util

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	//定义日志级别
	UNKNOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARING
	ERROR
	FATAL
)

type LogLevel uint16

//向终端写日志
type Logger struct {
	Level LogLevel
}

func (l Logger) enable(logLevel LogLevel) bool {
	return logLevel >= l.Level
}

func parseLogLevel(levelStr string) (LogLevel, error) {
	levelStr = strings.ToLower(levelStr)

	switch levelStr {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "waring":
		return WARING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil

	default:
		err := errors.New("没有此级别的日志")
		return UNKNOWN, err

	}
}
func NewLog(levelStr string) Logger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return Logger{
		Level: level,
	}
}

func (l Logger) Debug(msg string) {
	if l.enable(DEBUG) {
		now := time.Now()
		fmt.Printf("[%s] [DEBUG] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}
func (l Logger) Trace(msg string) {
	if l.enable(TRACE) {
		now := time.Now()
		fmt.Printf("[%s] [TRACE] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}
func (l Logger) Info(msg string) {
	if l.enable(INFO) {
		now := time.Now()
		fmt.Printf("[%s] [INFO] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}
func (l Logger) Warning(msg string) {
	if l.enable(WARING) {
		now := time.Now()
		fmt.Printf("[%s] [WARING] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}
func (l Logger) Error(msg string) {
	if l.enable(ERROR){
		now := time.Now()
		fmt.Printf("[%s] [ERROR] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}
func (l Logger) Fatal(msg string) {
	if l.enable(FATAL){
		now := time.Now()
		fmt.Printf("[%s] [FATAL] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}
