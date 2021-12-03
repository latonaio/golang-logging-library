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

func isJsonString(s interface{}) (b bool) {
	defer func() {
		if err := recover(); err != nil {
			b = false
		}
	}()
	val := reflect.ValueOf(s)
	for i := 0; i < val.Type().NumField(); i++ {
		if val.Type().Field(i).Tag.Get("json") != "" {
			return true
		}
	}
	return false
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
