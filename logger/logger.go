package logger

import (
	"fmt"
	"os"
	"time"
)

type Logger struct {
	Log map[string]interface{}
}

func NewLogger() *Logger {

	return &Logger{
		Log: map[string]interface{}{},
	}
}

// 文字列を引数に渡した場合は文字列を表示、JSONに対応したマップや構造体を引数に渡した場合はJSONを表示
func (*Logger) Fatal(msg interface{}, format ...interface{}) {
	log(msg, "FATAL", format)
	panic("Fatal error")
}

// 文字列を引数に渡した場合は文字列を表示、JSONに対応したマップや構造体を引数に渡した場合はJSONを表示
func (*Logger) Error(msg interface{}, format ...interface{}) {
	log(msg, "ERROR", format)
}

// 文字列を引数に渡した場合は文字列を表示、JSONに対応したマップや構造体を引数に渡した場合はJSONを表示
func (*Logger) Warn(msg interface{}, format ...interface{}) {
	log(msg, "WARN", format)
}

// 文字列を引数に渡した場合は文字列を表示、JSONに対応したマップや構造体を引数に渡した場合はJSONを表示
func (*Logger) Info(msg interface{}, format ...interface{}) {
	log(msg, "INFO", format)
}

// 文字列を引数に渡した場合は文字列を表示、JSONに対応したマップや構造体を引数に渡した場合はJSONを表示
func (*Logger) Debug(msg interface{}, format ...interface{}) {
	log(msg, "DEBUG", format)
}

func log(msg interface{}, logLevel string, variableStr []interface{}) {
	output := map[string]interface{}{
		"level":    logLevel,
		"time":     time.Now(),
		"cursor":   createCursor(),
		"function": createFunctionName(),
	}
	defer fin(output)

	// printf系の処理
	typedMsg, ok := msg.(string)
	if ok {
		output["message"] = fmt.Sprintf(typedMsg, variableStr...)
		return
	}

	// errorの出力
	_, ok = msg.(error)
	if ok {
		output["message"] = fmt.Sprintf("%+v", msg)
		return
	}

	// jsonに変換できる場合の処理
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
