package Translation

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mahdic200/weava/Utils/Constants"
)

type initializer struct {
	lang     string
	env_name string
	path     string
	object   structure
}

type structure struct {
	Fields   map[string]string `json:"fields"`
	Messages map[string]string `json:"messages"`
}

func (i *initializer) setLang(lang string, env_name string) error {
	if lang == "" {
		return fmt.Errorf("translationProvider : '%s' environment variable must be set and non-empty", env_name)
	}
	i.lang = lang
	return nil
}

func (i *initializer) setPath(path []string) {
	i.path = filepath.Join(Constants.BASE_DIR, filepath.Join(path...), i.lang+".json")
}

func (i *initializer) loadFile() error {
	data, err := os.ReadFile(i.path)
	if err != nil {
		return fmt.Errorf("translationProvider : %s", err.Error())
	}

	var translationObject structure
	if err := json.Unmarshal(data, &translationObject); err != nil {
		return fmt.Errorf("translationProvider : Invalid json format : at file %s : %s", i.path, err.Error())
	}
	i.object = translationObject
	return nil
}

func (i *initializer) validateObject() error {
	if i.object.Fields == nil {
		return fmt.Errorf("translationProvider : translation file must have 'fields' key")
	}
	if i.object.Messages == nil {
		return fmt.Errorf("translationProvider : translation file must have 'messages' key")
	}
	return nil
}

func GetTranslation(lang string, env_name string, path []string) (fields map[string]string, messages map[string]string, err error) {
	init := initializer{lang: lang, env_name: env_name}
	if err := init.setLang(lang, env_name); err != nil {
		return nil, nil, err
	}
	init.setPath(path)
	if err := init.loadFile(); err != nil {
		return nil, nil, err
	}
	if err := init.validateObject(); err != nil {
		return nil, nil, err
	}
	fields, messages, err = init.object.Fields, init.object.Messages, nil
	return
}
