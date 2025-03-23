package Providers

import (
	"fmt"

	"github.com/mahdic200/weava/Config"
)

func ErrorProvider(err error) error {
	message := err.Error()
	if !Config.APP_DEBUG {
		message = "Internal server error"
	}
	return fmt.Errorf("%s", message)
}
