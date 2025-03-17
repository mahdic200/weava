package ValidationProvider

import (
	"os"
	"testing"

	"github.com/mahdic200/weava/Providers/Validation"
)

func TestGetValLan(t *testing.T) {
    os.Setenv("VALIDATION_LANG", "")
    _, err := Validation.GetValidationLangKey()
    if err == nil {
        t.Error("GetValidationLangKey must give errors when env VALIDATION_LANG is empty !")
    }
}
