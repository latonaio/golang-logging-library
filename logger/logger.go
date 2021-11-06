package logger

import (
	"encoding/json"
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
	m := map[string]interface{}{}

	_, isString := msg.(string)
	if isString && len(variableStr) == 0 {
		m["message"] = msg
	} else if isString && len(variableStr) > 0 {

		msg = fmt.Sprintf(msg.(string), variableStr...)

		m["message"] = msg
	} else if isJsonString(msg) {
		// 渡されたデータを一度JSONに変換した後、マップに変換することで、構造体で定義されたJSONであってもマップとして扱う
		processingJsonData, _ := json.Marshal(msg)
		jsonObj := loadJson(processingJsonData)

		for k, v := range jsonObj {
			m[k] = v
		}
	} else {
		panic("this is not string or json")
	}

	m["level"] = logLevel
	m["time"] = time.Now()
	m["cursor"] = createCursor()

	switch logLevel {
	case "FATAL":
		fmt.Fprintln(os.Stderr, jsonParse(m))
	case "ERROR":
		fmt.Fprintln(os.Stderr, jsonParse(m))
	case "WARN":
		fmt.Fprintln(os.Stderr, jsonParse(m))
	case "INFO":
		fmt.Fprintln(os.Stdout, jsonParse(m))
	case "DEBUG":
		fmt.Fprintln(os.Stdout, jsonParse(m))
	}
}
