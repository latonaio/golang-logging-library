package test

import (
	"fmt"
	"testing"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

func Test_Struct(t *testing.T) {
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

func Test_Error(t *testing.T) {
	l := logger.NewLogger()
	err := fmt.Errorf("testError")
	err = xerrors.Errorf("WRAP ERROR : %w", err)

	l.Debug(err)
}

func Test_StructNoJsonTag(t *testing.T) {
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

func Test_StructArray(t *testing.T) {
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

func Test_StringArray(t *testing.T) {
	l := logger.NewLogger()
	strArr := []string{
		"test",
		"help",
	}

	l.Debug(strArr)
}

func Test_ErrorArray(t *testing.T) {
	// うまく動作しないが、エラー型の配列を渡すことは考えられないため放置
	l := logger.NewLogger()
	strArr := []error{
		fmt.Errorf("test"),
		fmt.Errorf("hello"),
	}

	l.Debug(strArr)
}

func Test_StructPointer(t *testing.T) {
	l := logger.NewLogger()
	jstr := &struct {
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

func Test_Option(t *testing.T) {
	l := logger.NewLogger()
	jstr := &struct {
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

	l.AddOption(map[string]interface{}{"test": 78}).Debug(jstr)
}
