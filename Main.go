package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bbemis017/ApartmentNotifier/notifications"
)

var g_notifier notifications.Notifier

func init() {
	fmt.Println("Init Lambda")

	var err error
	g_notifier, err = notifications.New()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Lambda Initialized")
}

func hello() (string, error) {
	fmt.Println("Logging the handler")

	g_notifier.Send()

	return "Hello ƛ! whatsuppp", nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(hello)
}
