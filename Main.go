package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bbemis017/ApartmentNotifier/datastore"
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
	process()
}

func hello() (string, error) {
	fmt.Println("Logging the handler")

	process()

	g_notifier.Send(notifications.NotifierContent{Unit: "5G"})

	return "Hello ƛ! whatsuppp", nil
}

func process() {
	timestamp := time.Now().Format(time.RFC3339)

	job := scrapeit.NewJob(20, true)
	job.Start()
	rawData, _ := job.AwaitResult()

	csvStore := datastore.NewCSVStore("apartments.csv")
	if csvStore.Length < 1 {
		datastore.WriteHeader(&csvStore)
	}

	log.Println("Write Apartment data")
	for _, val := range rawData["Apartments"].([]interface{}) {
		unit, _ := datastore.NewUnit(val.(map[string]interface{}), "Ravenswood Terrace", "1801 W Argyle St, Chicago, IL 60640", timestamp)

		unit.Save(&csvStore)
	}

	log.Println("Done")
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(hello)
}
