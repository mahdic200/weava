package Utils

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/mahdic200/weava/Config"
	"github.com/mahdic200/weava/Utils/Constants"
	"github.com/mahdic200/weava/Utils/File"
)

func PathToUrl(path string) string {
	sliced := strings.Split(path, string(filepath.Separator))
	url := strings.Join(sliced, "/")
	return url
}

func PathToHttpUrl(path string) string {
	url := strings.Join([]string{Config.APP_BASEURL, PathToUrl(path)}, "/")
	return url
}

func UserDefaultImage() string {
	dirname := "user_default_profile"
	path := Constants.PUBLIC_DIR + "/" + dirname
	entries, err := os.ReadDir(path)
	if err != nil {
		return ""
	}
	if len(entries) == 0 {
		return ""
	}
	return dirname + "/" + entries[0].Name()
}

func ImageUrlOrDefault(relative_path string) string {
	abs_path := File.PublicPath(relative_path)
	if !File.Exists(abs_path) {
		return PathToHttpUrl(UserDefaultImage())
	}
	return PathToHttpUrl(relative_path)
}
