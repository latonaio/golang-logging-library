package logger

import (
	"fmt"
	"os"
	"time"
)

type Logger struct {
	Log           map[string]interface{}
	headerInfoMsg map[string]interface{}
}

func NewLogger() *Logger {
	return &Logger{
		Log: map[string]interface{}{},
	}
}

func (l *Logger) AddHeaderInfo(msg map[string]interface{}) {
	if l.headerInfoMsg == nil {
		l.headerInfoMsg = make(map[string]interface{}, len(msg))
	}
	for k, v := range msg {
		switch k {
		case "level", "time", "cursor", "function", "message":
			continue
		}
		l.headerInfoMsg[k] = v
	}
}

// 文字列を引数に渡した場合は文字列を表示、JSONに対応したマップや構造体を引数に渡した場合はJSONを表示
func (l *Logger) Fatal(msg interface{}, format ...interface{}) {
	l.log(msg, "FATAL", format)
	panic("Fatal error")
}

// 文字列を引数に渡した場合は文字列を表示、JSONに対応したマップや構造体を引数に渡した場合はJSONを表示
func (l *Logger) Error(msg interface{}, format ...interface{}) {
	l.log(msg, "ERROR", format)
}

// 文字列を引数に渡した場合は文字列を表示、JSONに対応したマップや構造体を引数に渡した場合はJSONを表示
func (l *Logger) Warn(msg interface{}, format ...interface{}) {
	l.log(msg, "WARN", format)
}

// 文字列を引数に渡した場合は文字列を表示、JSONに対応したマップや構造体を引数に渡した場合はJSONを表示
func (l *Logger) Info(msg interface{}, format ...interface{}) {
	l.log(msg, "INFO", format)
}

// 文字列を引数に渡した場合は文字列を表示、JSONに対応したマップや構造体を引数に渡した場合はJSONを表示
func (l *Logger) Debug(msg interface{}, format ...interface{}) {
	l.log(msg, "DEBUG", format)
}

func (l *Logger) log(msg interface{}, logLevel string, variableStr []interface{}) {
	output := map[string]interface{}{
		"level":    logLevel,
		"time":     time.Now().Format(time.RFC3339),
		"cursor":   createCursor(),
		"function": createFunctionName(),
	}
	defer fin(output)

	// printf系の処理
	typedMsg, ok := msg.(string)
	if ok {
		output["message"] = fmt.Sprintf(typedMsg, variableStr...)
		if len(l.headerInfoMsg) > 0 {
			output["information"] = l.headerInfoMsg
		}
		return
	}

	// errorの出力
	_, ok = msg.(error)
	if ok {
		output["message"] = fmt.Sprintf("%+v", msg)
		if len(l.headerInfoMsg) > 0 {
			output["information"] = l.headerInfoMsg
		}
		return
	}
	// jsonに変換できる場合の処理
	if len(l.headerInfoMsg) > 0 {
		for k, v := range l.headerInfoMsg {
			output[k] = v
		}
	}
	output["message"] = msg
}
func fin(msg map[string]interface{}) {
	switch msg["level"] {
	case "FATAL", "ERROR", "WARN":
		fmt.Fprintln(os.Stderr, jsonParse(msg))
	case "INFO", "DEBUG":
		fmt.Fprintln(os.Stdout, jsonParse(msg))
	}
}
