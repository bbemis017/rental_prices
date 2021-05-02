package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bbemis017/ApartmentNotifier/notifications"
)

var g_notifier notifications.Notifier

func init() {
	fmt.Println("Init Lambda")

	g_notifier, _ = notifications.New()
	fmt.Println(g_notifier)
}

func hello() (string, error) {
	fmt.Println("Logging the handler")
	fmt.Println(g_notifier)
	g_notifier.Send()

	return "Hello ƛ! whatsuppp", nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(hello)
}
