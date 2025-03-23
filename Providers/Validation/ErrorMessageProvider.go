package Validation

import (
	"fmt"
	"strings"

	"github.com/mahdic200/weava/Config"
	"github.com/mahdic200/weava/Providers/Translation"
)

func ErrorMessageProvider(field_name string, rule_name string, otherKeys map[string]string) (message string, err error) {
	fields, messages, err := Translation.GetTranslation(Config.VALIDATION_LANG, "VALIDATION_LANG", []string{"Lang", "Validation"})
	if err != nil {
		return message, err
	}
	translatedField := fields[field_name]
	if translatedField == "" {
		translatedField = field_name
	}
	translatedMessage := messages[string(rule_name)]
	if translatedMessage == "" {
		return "", fmt.Errorf("error : no message is provided for this rule")
	}
	translatedMessage = strings.ReplaceAll(translatedMessage, "{field}", translatedField)
	for key, value := range otherKeys {
		translatedMessage = strings.ReplaceAll(translatedMessage, key, value)
	}
	return translatedMessage, nil
}
