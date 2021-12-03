package test

import (
	"fmt"
	"testing"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

func Test_a1(t *testing.T) {
	l := logger.NewLogger()
	jstr := struct {
		K1 int         "json:\"Key1\""
		K2 string      "json:\"Key2\""
		K3 string      "json:\"Key 3\""
		K4 interface{} "json:\"Key 4,omitempty\""
		K5 interface{} "json:\"Key 5\""
	}{
		K1: 1,
		K2: "test",
		K5: nil,
	}

	l.Debug(jstr)
}

func Test_a2(t *testing.T) {
	l := logger.NewLogger()
	err := fmt.Errorf("testError")
	err = xerrors.Errorf("WRAP ERROR : %w", err)

	l.Debug(err)
}
