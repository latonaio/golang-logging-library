package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
)

type Logger struct {
	Message interface{} `json:"message"`
	Level   string      `json:"level"`
	Cursor  string      `json:"cursor"`
	Time    interface{} `json:"time"`
	Tag     interface{} `json:"tag"`
}

//
func NewLogger() *Logger {

	return &Logger{
		Message: "default",
		Level:   "default",
		Cursor:  "default",
		Time:    "default",
		Tag:     nil,
	}
}

// 第二引数はオプション, 必要ない場合はnilに設定してください。
func (l *Logger) Fatal(msg interface{}, option interface{}) {
	l.log(msg, "FATAL", option)
	panic("Fatal error")
}

// 第二引数はオプション, 必要ない場合はnilに設定してください。
func (l *Logger) Error(msg interface{}, option interface{}) {
	l.log(msg, "ERROR", option)
}

// 第二引数はオプション, 必要ない場合はnilに設定してください。
func (l *Logger) Warn(msg interface{}, option interface{}) {
	l.log(msg, "WARN", option)
}

// 第二引数はオプション, 必要ない場合はnilに設定してください。
func (l *Logger) Info(msg interface{}, option interface{}) {
	l.log(msg, "INFO", option)
}

// 第二引数はオプション, 必要ない場合はnilに設定してください。
func (l *Logger) Debug(msg interface{}, option interface{}) {
	l.log(msg, "DEBUG", option)
}

// ログ表示
func (l *Logger) log(msg interface{}, logLevel string, option interface{}) {
	_, isTag := option.(string)
	if isTag {
		l.Tag = option
	} else {
		l.Tag = nil
	}

	l.Message = msg
	l.Level = logLevel
	l.Cursor = cursor()
	l.Time = time.Now()

	switch logLevel {
	case "FATAL":
		fmt.Fprintln(os.Stderr, l.parse())
	case "ERROR":
		fmt.Fprintln(os.Stderr, l.parse())
	case "WARN":
		fmt.Fprintln(os.Stderr, l.parse())
	case "INFO":
		fmt.Fprintln(os.Stdout, l.parse())
	case "DEBUG":
		fmt.Fprintln(os.Stdout, l.parse())
	}
}

// 行番号の表示
func cursor() string {
	_, file, line, _ := runtime.Caller(3)
	return file + "#L" + strconv.Itoa(line)
}

// jsonへパース
func (l *Logger) parse() string {
	result, err := json.Marshal(l)
	if err != nil {
		panic(err)
	}
	return string(result)
}
