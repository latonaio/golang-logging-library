package logger

import (
	"encoding/json"
	"reflect"
	"runtime"
	"strconv"
)

func createCursor() string {
	_, file, line, _ := runtime.Caller(3)
	return file + "#L" + strconv.Itoa(line)
}
func createFunctionName() string {
	pt, _, _, _ := runtime.Caller(3)
	return runtime.FuncForPC(pt).Name()
}

func isStructure(s interface{}) (b bool) {
	defer func() {
		if err := recover(); err != nil {
			b = false
		}
	}()

	val := reflect.ValueOf(s)
	for val.Type().Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Type().Kind() == reflect.Struct {
		return true
	}
	return val.Type().Kind() == reflect.Slice
}

func loadJson(byteArray []byte) map[string]interface{} {
	var jsonObj map[string]interface{}
	_ = json.Unmarshal(byteArray, &jsonObj)
	return jsonObj
}

func jsonParse(l map[string]interface{}) string {
	result, err := json.Marshal(l)
	if err != nil {
		panic(err)
	}
	return string(result)
}
