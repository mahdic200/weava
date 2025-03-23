package Validation

import (
	"fmt"
	"strings"
)

func ErrorMessageProvider(field_name string, rule_name string, otherKeys map[string]string) (message string, err error) {
	fields, messages, err := TranslationProvider()
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
