package notifications

import "fmt"

type LogMessage struct{}

func NewLogMessage() (Notifier, error) {
	return LogMessage{}, nil
}

func (logMessage LogMessage) Send(content NotifierContent) error {
	fmt.Println("Notifier " + content.toString())
	return nil
}
