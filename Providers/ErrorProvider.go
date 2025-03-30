package Providers

import (
	"github.com/mahdic200/weava/Config"
)

func ErrorProvider(err error) string {
	message := err.Error()
	if !Config.APP_DEBUG {
		message = "Internal server error"
	}
	return message
}
