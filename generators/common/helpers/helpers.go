package helpers

import (
	"path"
	"path/filepath"
	"runtime"
)

func getCurrentPath() string {
	_, filename, _, _ := runtime.Caller(1)

	return path.Dir(filename)
}

// BASE_PATH is the base path to access the data informations
var BASE_PATH, _ = filepath.Abs(getCurrentPath() + `../../../../`)

func DecodeString(something []interface{}) string {
	if len(something) == 0 {
		return ""
	}
	return something[0].(string)
}
func DecodeUint64(something []interface{}) uint64 {
	if len(something) == 0 {
		return 0
	}
	return uint64(something[0].(uint8))
}
