package Constants

import (
	"os"
	"path/filepath"
)

func GetBaseDir() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	return dir
}

var BASE_DIR string = GetBaseDir()

var PUBLIC_DIR string = filepath.Join(BASE_DIR, "public")

var UPLOADS_PATH string = filepath.Join(PUBLIC_DIR, "uploads")

var VERSION string = "0.0.5"
