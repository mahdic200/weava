package Response

import (
	"fmt"
	"strings"

	"github.com/mahdic200/weava/Config"
	"github.com/mahdic200/weava/Providers/Translation"
)

func Message(field string, key string) (message string, err error) {
	fields, messages, err := Translation.GetTranslation(Config.RESMSG_LANG, "RESMSG_LANG", []string{"Lang", "ResponseMessage"})
	if err != nil {
		return "", err
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
