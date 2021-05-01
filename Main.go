package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

func hello() (string, error) {
	fmt.Println("Logging the handler")

	test_mail()

	return "Hello ƛ! whatsuppp", nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(hello)
}

type ClientMessage struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

var toEmail string
var subject string
var emailClient *ses.SES

func init() {
	fmt.Println("init")
	toEmail = os.Getenv("TO_EMAIL")
	subject = os.Getenv("SUBJECT")

	if len(subject) < 0 {
		subject = "Message from website"
	}

	emailClient = ses.New(session.New(), aws.NewConfig().WithRegion("us-east-1"))
}

func test_mail() {
	fmt.Println("Test send mail")

	emailParams := &ses.SendEmailInput{
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Data: aws.String("message text"),
				},
			},
			Subject: &ses.Content{
				Data: aws.String(subject),
			},
		},
		Destination: &ses.Destination{
			ToAddresses: []*string{aws.String(toEmail)},
		},
		Source: aws.String(toEmail),
	}
	fmt.Println(emailParams)

	result, err := emailClient.SendEmail(emailParams)

	fmt.Println("Response")
	fmt.Println(result)

	if err != nil {
		fmt.Println("Error")
	}
}
