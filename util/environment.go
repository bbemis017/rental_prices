package util

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	ENV_AWS_REGION      = "AWS_REGION"
	ENV_AWS_S3_BUCKET   = "AWS_S3_BUCKET"
	ENV_NOTIFY_TYPE     = "NOTIFY_TYPE"
	ENV_EMAIL_RECIPIENT = "EMAIL_RECIPIENT"
	ENV_EMAIL_SUBJECT   = "EMAIL_SUBJECT"
	ENV_APARTMENTS_CSV  = "APARTMENTS_CSV"

	ENV_SCRAPEIT_NET_HOST  = "SCRAPEIT_NET_HOST"
	ENV_SCRAPEIT_NET_KEY   = "SCRAPEIT_NET_KEY"
	ENV_SCRAPEIT_NET_CACHE = "SCRAPEIT_NET_CACHE"

	ENV_LAMBDA_MODE = "LAMBDA_MODE"
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

func GetEnvBoolOrFail(key string) bool {
	strValue := GetEnvOrFail(key)
	value, err := strconv.ParseBool(strValue)
	if err != nil {
		log.Fatalf("%s is not a boolean", key)
	}
	return value
}

func GetEnvBoolOrDefault(key string, defaultValue bool) bool {
	strVal := GetEnvOrDefault(key, strconv.FormatBool(defaultValue))
	val, err := strconv.ParseBool(strVal)
	if err != nil {
		log.Fatalf("%s is not a bollean", key)
	}
	return val
}

func GetEnvOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
