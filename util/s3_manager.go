package util

import (
	"log"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func SaveToS3(bucket string, filegroup string, data string) {

	// Create a filename with the file group and timestamp
	timestamp := time.Now().Format(time.RFC3339)
	filename := filegroup + "_" + timestamp + ".csv"

	// create a reader from data in memory
	reader := strings.NewReader(data)

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(GetEnvOrFail(ENV_AWS_REGION))},
	)
	uploader := s3manager.NewUploader(sess)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		// here you pass your reader
		// the aws sdk will manage all the memory and file reading for you
		Body: reader,
	})
	if err != nil {
		log.Fatalf("Unable to upload %q to %q, %v", filename, bucket, err)
	}

	log.Printf("Successfully uploaded %q to %q\n", filename, bucket)
}
