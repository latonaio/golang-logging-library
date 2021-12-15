package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Logger interface {
	Fatal(Msg interface{}, args ...interface{})
	Error(Msg interface{}, args ...interface{})
	Warn(Msg interface{}, args ...interface{})
	Info(Msg interface{}, args ...interface{})
	Debug(Msg interface{}, args ...interface{})

	AddOption(map[string]interface{}) Logger
}

func NewLogger() Logger {
	return &logger{}
}

type logger struct{}

func (l *logger) Fatal(Msg interface{}, args ...interface{}) {
	newTempPrinter().Fatal(Msg, args...)
}

func (l *logger) Error(Msg interface{}, args ...interface{}) {
	newTempPrinter().Error(Msg, args...)
}

func (l *logger) Warn(Msg interface{}, args ...interface{}) {
	newTempPrinter().Warn(Msg, args...)
}

func (l *logger) Info(Msg interface{}, args ...interface{}) {
	newTempPrinter().Info(Msg, args...)
}

func (l *logger) Debug(Msg interface{}, args ...interface{}) {
	newTempPrinter().Debug(Msg, args...)
}

func (l *logger) AddOption(fields map[string]interface{}) Logger {
	return &tempPrinter{
		msg: fields,
	}
}

const (
	MSG  = "message"
	LVL  = "level"
	TIME = "time"
	CUR  = "cursor"
	FUNC = "function"
)

type tempPrinter struct {
	msg
}
type msg map[string]interface{}

func newTempPrinter() *tempPrinter {
	return &tempPrinter{
		msg: make(msg, 0),
	}
}

// // 文字列を引数に渡した場合は文字列を表示、JSONに対応したマップや構造体を引数に渡した場合はJSONを表示
func (p *tempPrinter) Fatal(msg interface{}, format ...interface{}) {
	lvl := "FATAL"
	p.build(msg, lvl, format)
	p.printout(lvl)
	panic("Fatal error")
}

// 文字列を引数に渡した場合は文字列を表示、JSONに対応したマップや構造体を引数に渡した場合はJSONを表示
func (p *tempPrinter) Error(msg interface{}, format ...interface{}) {
	lvl := "ERROR"
	p.build(msg, lvl, format)
	p.printout(lvl)
}

// 文字列を引数に渡した場合は文字列を表示、JSONに対応したマップや構造体を引数に渡した場合はJSONを表示
func (p *tempPrinter) Warn(msg interface{}, format ...interface{}) {
	lvl := "WARN"
	p.build(msg, lvl, format)
	p.printout(lvl)
}

// 文字列を引数に渡した場合は文字列を表示、JSONに対応したマップや構造体を引数に渡した場合はJSONを表示
func (p *tempPrinter) Info(msg interface{}, format ...interface{}) {
	lvl := "INFO"
	p.build(msg, lvl, format)
	p.printout(lvl)
}

// 文字列を引数に渡した場合は文字列を表示、JSONに対応したマップや構造体を引数に渡した場合はJSONを表示
func (p *tempPrinter) Debug(msg interface{}, format ...interface{}) {
	lvl := "DEBUG"
	p.build(msg, lvl, format)
	p.printout(lvl)
}

func (l *tempPrinter) AddOption(fields map[string]interface{}) Logger {
	if l == nil || len(l.msg) == 0 {
		l = newTempPrinter()
	}
	for k, v := range fields {
		l.msg[k] = v
	}

	return l
}

func (p *tempPrinter) build(msg interface{}, logLevel string, variableStr []interface{}) {
	p.setMetaData(logLevel)
	switch m := msg.(type) {
	case string:
		p.msg[MSG] = fmt.Sprintf(m, variableStr...)
	case error:
		p.msg[MSG] = fmt.Sprintf("%+v", m)
	default:
		p.msg[MSG] = m
	}
}

func (p *tempPrinter) setMetaData(logLevel string) {
	p.msg[LVL] = logLevel
	p.msg[TIME] = time.Now()
	p.msg[CUR] = createCursor()
	p.msg[FUNC] = createFunctionName()
}

func (p *tempPrinter) printout(logLevel string) {
	switch logLevel {
	case "FATAL", "ERROR", "WARN":
		fmt.Fprintln(os.Stderr, p)
	case "INFO", "DEBUG":
		fmt.Fprintln(os.Stdout, p)
	}
}

func (p *tempPrinter) String() string {
	if p == nil || len(p.msg) == 0 {
		return "[FATAL?] logger is nil ?"
	}
	result, err := json.Marshal(p.msg)
	if err != nil {
		return fmt.Sprintf("[FATAL?] log parse error raw message: %+v", p.msg[MSG])
	}
	return string(result)
}
