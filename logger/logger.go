package logger

import (
	"fmt"
	"os"
	"strings"
)

// ログのフォーマットの指定
const T = "{TIME}"   // {TIME} : 日時
const LV = "{LEVEL}" // {LEVEL} : ログレベル
const M = "{MSG}"    // {MSG} : メッセージ詳細
const LG = "{LOG}"   // {LOG} : ログメッセージ msgFormatのこと
const F = "{FILE}"   // {FILE} : ログ出力が呼ばれたファイ
const LN = "{LINE}"  // {LINE} : ログ出力が呼ばれた行番号

// msgFormat ログメッセージのフォーマット
const msgFmt = `{[{LEVEL}]{FILE}:L{LINE}\n\t{MSG}}`

// logFormat ログ出力全体のフォーマット
const logFmt = `{"log":"{LOG}","time":{TIME}}"`

// timeFormat ログ出力に使用する時間のフォーマット
const timeFmt = "2006/01/02 15:04:05.99"

// Logger
type Logger struct {
	option []func(msg interface{})
	level  int
	t      *tmp
}

type tmp struct {
	log    string
	msg    string
	rawMsg interface{}
	lvl    int
}

// NewLogger ログ用構造体を新しく作成する
// level : ログ出力のレベルを指定する
// 0 : 出力なし
// 1 : Fatal
// 2 : Error
// 3 : Info
// 4 : Debug
func NewLogger(level uint) *Logger {
	if level > 4 {
		level = 3
	}
	return &Logger{
		option: make([]func(interface{}), 0),
		level:  int(level),
		t:      nil,
	}
}

var msgArgFuncs []func(*tmp, string) string
var logArgFuncs []func(*tmp, string) string

func init() {
	setInitVars(&msgArgFuncs, msgFmt)
	setInitVars(&logArgFuncs, logFmt)
}

func setInitVars(funcQueue *[]func(*tmp, string) string, format string) {
	t := strings.Index(format, T)
	lv := strings.Index(format, LV)
	f := strings.Index(format, F)
	ln := strings.Index(format, LN)
	lg := strings.Index(format, LG)
	m := strings.Index(format, M)

	order := []int{t, lv, f, ln, lg, m}
	queue := make([]func(*tmp, string) string, 0, len(order))

	for i := 0; i < len(order); i++ {
		o := order[i]
		if o < 0 {
			continue
		}
		switch o {
		case t:
			queue = append(queue, rTIME)
		case lv:
			queue = append(queue, rLEVEL)
		case f:
			queue = append(queue, rFILE)
		case ln:
			queue = append(queue, rLINE)
		case m:
			queue = append(queue, rMSG)
		case lg:
			queue = append(queue, rLOG)
		}
	}
	*funcQueue = queue
}

func (l *Logger) Fatal(msg interface{}) {
	l.prepare(msg, 1)
	fmt.Fprintln(os.Stderr, l.t.log)
	l.suf()

	panic("Fatal error")
}
func (l *Logger) Error(msg interface{}) {
	if l.level < 2 {
		return
	}
	l.prepare(msg, 2)
	fmt.Fprintln(os.Stderr, l.t.log)
	l.suf()
}
func (l *Logger) Info(msg interface{}) {
	if l.level < 3 {
		return
	}
	l.prepare(msg, 3)
	fmt.Fprintln(os.Stdout, l.t.log)
	l.suf()
}
func (l *Logger) Debug(msg interface{}) {
	if l.level < 4 {
		return
	}
	l.prepare(msg, 4)
	fmt.Fprintln(os.Stdout, l.t.log)
	l.suf()
}
func (l *Logger) AddOption(f func(interface{})) {
	l.option = append(l.option, f)
}

func (l *Logger) prepare(msg interface{}, lvl int) {
	l.t = &tmp{
		log:    logFmt,
		msg:    msgFmt,
		rawMsg: msg,
		lvl:    lvl,
	}

	l.buildMsg()
	l.buildLog()
}

func (l *Logger) buildMsg() {
	for i := range msgArgFuncs {
		l.t.msg = msgArgFuncs[i](l.t, l.t.msg)
	}
}

func (l *Logger) buildLog() {
	for i := range logArgFuncs {
		l.t.log = logArgFuncs[i](l.t, l.t.log)
	}
}

func (l *Logger) suf() {
	for i := range l.option {
		l.option[i](l.t.log)
	}
}
