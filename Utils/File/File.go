package File

import (
	"os"
	"path/filepath"

	"github.com/mahdic200/weava/Utils/Constants"
)

func PublicPath(path string) string {
	frp := filepath.Join(Constants.PUBLIC_DIR, path)
	return frp
}

func Exists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	} else if info.IsDir() {
		return false
	}
	return true
}
