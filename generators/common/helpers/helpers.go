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
