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
	"github.com/bbemis017/ApartmentNotifier/util"
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

	process()

	g_notifier.Send(notifications.NotifierContent{Unit: "5G"})

	return "Hello ƛ! whatsuppp", nil
}

func process() {
	timestamp := time.Now().UTC().Format(time.RFC3339)

	csvStore := datastore.NewCSVStore([]string{"created_at", "complex", "unit_number", "price", "availability", "bedrooms", "baths", "address"})

	job := scrapeit.NewJob(20, true)
	job.Start()
	rawData, _ := job.AwaitResult()

	log.Println("Write Apartment data")
	for _, val := range rawData["Apartments"].([]interface{}) {
		unit, _ := datastore.NewUnit(val.(map[string]interface{}), "Ravenswood Terrace", "1801 W Argyle St, Chicago, IL 60640", timestamp)

		unit.Save(&csvStore)
	}

	s3Bucket := util.GetEnvOrDefault(util.ENV_AWS_S3_BUCKET, "NONE")
	if s3Bucket != "NONE" {
		util.SaveToS3(s3Bucket, "apartments", csvStore.Data)
	} else {
		log.Println("S3 Bucket not specified")
		log.Println("Logging Data to stdout")
		log.Println(csvStore.Data)
	}

	log.Println("Done")
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(hello)
}
