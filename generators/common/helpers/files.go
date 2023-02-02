package helpers

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
)

// BASE_PATH is the base path to access the data informations
var BASE_PATH, _ = filepath.Abs(getCurrentPath() + `../../../../`)

func getCurrentPath() string {
	_, filename, _, _ := runtime.Caller(1)

	return path.Dir(filename)
}

// CreateFile creates a file with the given path and returns the file object
func CreateFile(p string) error {
	if err := os.MkdirAll(p, 0770); err != nil {
		return err
	}
	return nil
}
