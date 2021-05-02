package util

import (
	"errors"
	"fmt"
	"os"
)

const (
	ENV_NOTIFY_TYPE = "NOTIFY_TYPE"
)

func GetEnvOrFail(key string) string {

	value := os.Getenv(key)
	if value == "" {
		err := errors.New(key + " is undefined")
		fmt.Println(err)
		os.Exit(1)
	}

	return value
}
