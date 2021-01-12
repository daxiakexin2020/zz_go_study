package util

import "fmt"

//向终端写日志

type Logger struct {

}

func NewLog() Logger {
	return Logger{}
}

func (l Logger) Debug(msg string)  {
	fmt.Println(msg)
}