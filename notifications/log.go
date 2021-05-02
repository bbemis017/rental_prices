package notifications

import "fmt"

type LogMessage struct {
	message string
}

func NewLogMessage(message string) (Notifier, error) {
	return LogMessage{message: message}, nil
}

func (logMessage LogMessage) Send() {
	fmt.Println(logMessage.message)
}
