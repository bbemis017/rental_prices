package notifications

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type EmailMessage struct {
	emailClient *ses.SES
	toEmail     string
	subject     string
}

func NewEmailMessage() (Notifier, error) {
	fmt.Println("Init Notifier")

	notifier := EmailMessage{
		toEmail: os.Getenv("TO_EMAIL"),
		subject: os.Getenv("SUBJECT"),
	}

	if len(notifier.subject) < 0 {
		notifier.subject = "Message from website"
	}

	notifier.emailClient = ses.New(session.New(), aws.NewConfig().WithRegion("us-east-1"))

	return notifier, nil
}

func (message EmailMessage) Send() {
	fmt.Println("Test send mail")

	emailParams := &ses.SendEmailInput{
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Data: aws.String("message text"),
				},
			},
			Subject: &ses.Content{
				Data: aws.String(message.subject),
			},
		},
		Destination: &ses.Destination{
			ToAddresses: []*string{aws.String(message.toEmail)},
		},
		Source: aws.String(message.toEmail),
	}
	fmt.Println(emailParams)

	result, err := message.emailClient.SendEmail(emailParams)

	fmt.Println("Response")
	fmt.Println(result)

	if err != nil {
		fmt.Println("Error")
	}
}
