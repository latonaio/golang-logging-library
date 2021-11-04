package main

import (
	"bitbucket.org/latonaio/golang-logging-library/logger"
)

func main() {
	l := logger.NewLogger()

	// normal
	l.Debug("test", nil)
	l.Debug("test", 111)
	l.Debug("test", "test")
	l.Debug("test", 111)

}
