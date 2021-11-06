package logger

import (
	"encoding/json"
	"runtime"
	"strconv"
)

func createCursor() string {
	_, file, line, _ := runtime.Caller(3)
	return file + "#L" + strconv.Itoa(line)
}

func isJsonString(s interface{}) bool {
	_, err := json.Marshal(s)

	return err == nil
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
