package File

import (
	"path/filepath"

	"github.com/mahdic200/weava/Utils/Constants"
)

func PublicPath(path string) string {
	frp := filepath.Join(Constants.PUBLIC_DIR, path)
	return frp
}
