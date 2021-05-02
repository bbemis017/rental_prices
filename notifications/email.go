package notifications

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/bbemis017/ApartmentNotifier/util"
)

type EmailMessage struct {
	emailClient *ses.SES
	toEmail     string
	subject     string
}

func NewEmailMessage() (Notifier, error) {

	notifier := EmailMessage{
		toEmail: util.GetEnvOrFail(util.ENV_EMAIL_RECIPIENT),
		subject: util.GetEnvOrDefault(util.ENV_EMAIL_SUBJECT, "ApartmentNotifier"),
	}

	notifier.emailClient = ses.New(
		session.New(),
		aws.NewConfig().WithRegion(util.GetEnvOrFail(util.ENV_AWS_REGION)),
	)

	return notifier, nil
}

func (message EmailMessage) Send(content NotifierContent) error {

	emailParams := &ses.SendEmailInput{
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Data: aws.String(content.toString()),
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

	_, err := message.emailClient.SendEmail(emailParams)
	return err
}
