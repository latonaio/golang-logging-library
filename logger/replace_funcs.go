package logger

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

func rTIME(d *tmp, format string) string {
	t := time.Now().Format(timeFmt)
	return strings.Replace(format, T, t, 1)
}

func rLEVEL(d *tmp, format string) string {
	lvlString := ""
	switch d.lvl {
	case 1:
		lvlString = "Fatal"
	case 2:
		lvlString = "Error"
	case 3:
		lvlString = "Info "
	case 4:
		lvlString = "Debug"
	default:
		panic("unknow log level")
	}
	return strings.Replace(format, LV, lvlString, 1)
}

func rMSG(d *tmp, format string) string {
	tmpFmt := strings.Replace(format, M, "%+v", 1)
	return fmt.Sprintf(tmpFmt, d.rawMsg)
}

func rLOG(d *tmp, format string) string {
	return strings.Replace(format, LG, d.msg, 1)
}

func rFILE(d *tmp, format string) string {
	_, f, _, ok := runtime.Caller(4)
	if !ok {
		panic("cannot get file info")
	}

	return strings.Replace(format, F, f, 1)
}

func rLINE(d *tmp, format string) string {
	_, _, li, ok := runtime.Caller(4)
	if !ok {
		panic("cannot get file info")
	}

	return strings.Replace(format, LN, fmt.Sprintf("%d", li), 1)
}
