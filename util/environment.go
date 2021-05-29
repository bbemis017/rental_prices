package util

import (
	"errors"
	"fmt"
	"os"
)

const (
	ENV_AWS_REGION      = "AWS_REGION"
	ENV_AWS_S3_BUCKET   = "AWS_S3_BUCKET"
	ENV_NOTIFY_TYPE     = "NOTIFY_TYPE"
	ENV_EMAIL_RECIPIENT = "EMAIL_RECIPIENT"
	ENV_EMAIL_SUBJECT   = "EMAIL_SUBJECT"
	ENV_APARTMENTS_CSV  = "APARTMENTS_CSV"

	ENV_SCRAPEIT_NET_HOST = "SCRAPEIT_NET_HOST"
	ENV_SCRAPEIT_NET_KEY  = "SCRAPEIT_NET_KEY"
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

func GetEnvOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
