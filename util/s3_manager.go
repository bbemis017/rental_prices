package util

import (
	"bytes"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func UploadFile(fileDir string, bucket string) {

	// Create a single AWS session (we can re use this if we're uploading many files)
	sess, sessionError := session.NewSession(&aws.Config{
		Region: aws.String(GetEnvOrFail(ENV_AWS_REGION)),
	})

	if sessionError != nil {
		log.Fatal(sessionError)
	}

	uploadErr := AddFileToS3(sess, fileDir, bucket)
	if uploadErr != nil {
		log.Fatal(uploadErr)
	}
}

// AddFileToS3 will upload a single file to S3, it will require a pre-built aws session
// and will set file info like content type and encryption on the uploaded file.
func AddFileToS3(s *session.Session, fileDir string, bucket string) error {

	// Open the file for use
	file, err := os.Open(fileDir)
	if err != nil {
		return err
	}
	defer file.Close()

	// Get file size and read the file content into a buffer
	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	// Config settings: this is where you choose the bucket, filename, content-type etc.
	// of the file you're uploading.
	_, err = s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(bucket),
		Key:                  aws.String(fileDir),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(size),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})
	return err
}
