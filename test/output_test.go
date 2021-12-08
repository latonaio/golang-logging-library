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

func Test_a3(t *testing.T) {
	l := logger.NewLogger()
	jstr := struct {
		K1 int
		K2 string
		K3 string
		K4 interface{}
		K5 interface{}
	}{
		K1: 1,
		K2: "test",
		K5: nil,
	}

	l.Debug(jstr)
}

func Test_a4(t *testing.T) {
	l := logger.NewLogger()
	jstr := []struct {
		K1 int         "json:\"Key1\""
		K2 string      "json:\"Key2\""
		K3 string      "json:\"Key 3\""
		K4 interface{} "json:\"Key 4,omitempty\""
		K5 interface{} "json:\"Key 5\""
	}{
		{
			K1: 1,
			K2: "test",
			K5: nil,
		},
		{
			K1: 2,
			K2: "test2",
			K5: "hello",
		},
	}

	l.Debug(jstr)
}

func Test_a5(t *testing.T) {
	l := logger.NewLogger()
	strArr := []string{
		"test",
		"help",
	}

	l.Debug(strArr)
}

func Test_a6(t *testing.T) {
	l := logger.NewLogger()
	strArr := []error{
		fmt.Errorf("test"),
		fmt.Errorf("hello"),
	}

	l.Debug(strArr)
}
