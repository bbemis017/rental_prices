package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bbemis017/ApartmentNotifier/notifications"
	"github.com/bbemis017/ApartmentNotifier/scrapeit"
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

	job := scrapeit.NewJob(15, true)
	fmt.Println(job.Start())

	fmt.Println(job.Status())

	g_notifier.Send(notifications.NotifierContent{Unit: "5G"})

	return "Hello ƛ! whatsuppp", nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(hello)
}
