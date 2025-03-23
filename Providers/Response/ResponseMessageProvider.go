package Response

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mahdic200/weava/Config"
	"github.com/mahdic200/weava/Utils/Constants"
)

func GetLangKey() (lang string, err error) {
	lang = Config.RESMSG_LANG
	if lang == "" {
		err = fmt.Errorf("resMsgProvider Error : 'RESMSG_LANG' key in env file must not be empty")
	}
	return
}

func GetTranslationPath() (path string, err error) {
	lang, err := GetLangKey()
	if err != nil {
		return "", err
	}
	path, err = filepath.Join(Constants.BASE_DIR, "Lang", "ResponseMessage", lang+".json"), nil
	return
}

func LoadTranslationFile() (fields map[string]string, messages map[string]string, err error) {
	path, err := GetTranslationPath()
	if err != nil {
		return nil, nil, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, fmt.Errorf("resMsgProvider Error : Failed to load language file : %w", err)
	}

	var translation map[string]map[string]string
	if err := json.Unmarshal(data, &translation); err != nil {
		return nil, nil, fmt.Errorf("resMsgProvider Error : Invalid json format : at file %s %w", path, err)
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
		return nil, nil, fmt.Errorf("resMsgProvider Error : Translation file must have 'fields' key")
	}
	if messages == nil {
		return nil, nil, fmt.Errorf("resMsgProvider Error : Translation file must have 'messages' key")
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

func Message(field string, key string) (message string, err error) {
	fields, messages, err := TranslationProvider()
	if err != nil {
		return message, err
	}
	translatedField := fields[field]
	if translatedField == "" {
		translatedField = field
	}
	translatedMessage := messages[key]
	if translatedMessage == "" {
		return "Operation completed (this is the default message)", fmt.Errorf("error : no message is provided for this ResponseMessage")
	}
	translatedMessage = strings.ReplaceAll(translatedMessage, "{field}", translatedField)
	return translatedMessage, nil
}
