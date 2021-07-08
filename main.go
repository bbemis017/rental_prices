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

	// disabling email notifications
	// g_notifier.Send(notifications.NotifierContent{Unit: "5G"})

	return "Hello ƛ! whatsuppp", nil
}

/* Calls a WebScraping Template, cleans data and appends metaData values to every record
 * @param templateId
 * @param metaData
 * @returns cleaned data in quoted csv format as a string
 * @returns TODO error if template could not be processed
 */
func process_complex(templateId int, metaData map[string]string) (string, error) {
	timestamp := util.FormatTimeStamp(time.Now())

	job := scrapeit.NewJob(28, util.GetEnvBoolOrFail(util.ENV_SCRAPEIT_NET_CACHE))
	job.Start()
	rawData, _ := job.AwaitResult()

	csvStr := ""

	header := []string{"created_at", "complex", "unit_number", "price", "availability", "bedrooms", "baths", "address"}
	for _, val := range rawData["apartments"].([]interface{}) {
		dataMap := val.(map[string]interface{})

		// static field values
		dataMap["created_at"] = timestamp

		addMetaData(dataMap, metaData)

		datastore.CleanDataMap(dataMap)

		csvStr += datastore.MapJsonToCsvString(header, dataMap)
	}

	return csvStr, nil
}

func addMetaData(record map[string]interface{}, metaData map[string]string) {
	for k, v := range metaData {
		record[k] = v
	}
}

func process() {
	log.Println("Process")

	ravenswoodTerrace, _ := process_complex(28, map[string]string{"address": "1801 W Argyle St, Chicago, IL 60640"})

	data := ravenswoodTerrace

	s3Bucket := util.GetEnvOrDefault(util.ENV_AWS_S3_BUCKET, "NONE")
	if s3Bucket != "NONE" {
		util.SaveToS3(s3Bucket, "apartments", data)
	} else {
		log.Println("S3 Bucket not specified")
		log.Println("Logging Data to stdout")
		log.Println(data)
	}

	log.Println("Done")
}

func main() {

	if util.GetEnvBoolOrDefault(util.ENV_LAMBDA_MODE, true) {
		// Make the handler available for Remote Procedure Call by AWS Lambda
		lambda.Start(hello)
	} else {
		hello()
	}
}
