package Validation

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mahdic200/weava/Config"
)

func GetValidationLangKey() (lang string, err error) {
	lang = Config.VALIDATION_LANG
	if lang == "" {
		err = fmt.Errorf("validationProvider Error : 'VALIDATION_LANG' key in env file must not be empty")
	}
	return
}

func GetTranslationPath() (path string, err error) {
	lang, err := GetValidationLangKey()
	if err != nil {
		return "", err
	}
	pwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("validationProvider Error : Failed to get the pwd")
	}
	path, err = filepath.Join(pwd, "Lang", "Validation", lang+".json"), nil
	return
}

func LoadTranslationFile() (fields map[string]string, messages map[string]string, err error) {
	path, err := GetTranslationPath()
	if err != nil {
		return nil, nil, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, fmt.Errorf("ValidationProvider Error : Failed to load language file : %w", err)
	}

	var translation map[string]map[string]string
	if err := json.Unmarshal(data, &translation); err != nil {
		return nil, nil, fmt.Errorf("ValidationProvider Error : Invalid json format : at file %s %w", path, err)
	}

	fields, messages, err = translation["fields"], translation["messages"], nil
	return
}

func VerifyTranslation() (fields map[string]string, messages map[string]string, err error) {
	fields, messages, err = LoadTranslationFile()
	if err != nil {
		return nil, nil, err
	}

	if fields == nil {
		return nil, nil, fmt.Errorf("validationProvider Error : Translation file must have 'fields' key")
	}
	if messages == nil {
		return nil, nil, fmt.Errorf("validationProvider Error : Translation file must have 'messages' key")
	}
	return
}

func TranslationProvider() (fields map[string]string, messages map[string]string, err error) {
	fields, messages, err = VerifyTranslation()
	if err != nil {
		return nil, nil, err
	}
	return
}
