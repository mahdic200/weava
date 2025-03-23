package FileService

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/mahdic200/weava/Utils"
	"github.com/mahdic200/weava/Utils/Constants"
)

// Define a regex to remove unwanted characters from filenames
var invalidChars = regexp.MustCompile(`[~'@#!\` + "`" + `]`)

type FileService struct {
	file           *multipart.FileHeader
	file_name      string
	file_extension string
	paths          []string
}

func GetFileExtension(filename string) string {
	sliced_name := strings.Split(filename, ".")
	fileExtension := sliced_name[len(sliced_name)-1]
	return fileExtension
}

// Sanitize filename to remove invalid characters
func sanitizeFileName(name string) string {
	return invalidChars.ReplaceAllString(name, "_")
}

func (fs *FileService) initFileName() {
	fs.file_name = fmt.Sprintf("%s.%s", Utils.StandardRandomString(32), fs.file_extension)
}

func (fs *FileService) SetFileName(name string) {
	fs.file_name = sanitizeFileName(name)
}

func (fs *FileService) relativeDirPath() string {
	return filepath.Join(fs.paths...)
}

func (fs *FileService) finalDirPath() string {
	return filepath.Join(Constants.PUBLIC_DIR, fs.relativeDirPath())
}

func (fs *FileService) relativeFilePath() string {
	return filepath.Join(fs.relativeDirPath(), fs.file_name)
}

func (fs *FileService) finalFilePath() string {
	return filepath.Join(fs.finalDirPath(), fs.file_name)
}

func (fs *FileService) GetRelativePath() string {
	return fs.relativeFilePath()
}

func (fs *FileService) GetFinalPath() string {
	return fs.finalFilePath()
}

func (fs *FileService) SaveToPublic(paths ...string) error {
	/* This is necessary , because functions related to handling file path work based on
	   this field */
	fs.paths = paths

	/* Getting the absolute path for initializing path */
	os.MkdirAll(fs.finalDirPath(), os.ModePerm)

	/* Now the path is initialized and new file could be placed in it safely without receiving any OS
	   errors such as : "no such a file or directory" */

	f, err := fs.file.Open()
	if err != nil {
		return err
	}

	/* Using defer function to close opened file in case of any unhandled panics */
	defer func() {
		f.Close()
	}()

	ff, osError := os.Create(fs.finalFilePath())

	/* Using defer function to close opened file in case of any unhandled panics */
	defer func() {
		ff.Close()
	}()

	if osError != nil {
		return fmt.Errorf("FileService Error , SaveToPublic Method : %s", err)
	}
	_, ioError := io.Copy(ff, f)
	if ioError != nil {
		return err
	}
	ff.Close()
	return nil
}

func New(file *multipart.FileHeader) FileService {
	fs := FileService{file: file, file_extension: GetFileExtension(file.Filename)}
	fs.initFileName()
	return fs
}
