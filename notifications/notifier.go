package notifications

import (
	"errors"

	"github.com/bbemis017/ApartmentNotifier/util"
)

type (
	NotifierContent struct {
		Unit string
	}

	Notifier interface {
		Send(content NotifierContent) error
	}
)

const (
	EMAIL_NOTIFY_TYPE = "EMAIL"
	LOG_NOTIFY_TYPE   = "LOG"
)

func New() (Notifier, error) {

	notify_type := util.GetEnvOrFail(util.ENV_NOTIFY_TYPE)
	switch notify_type {
	case EMAIL_NOTIFY_TYPE:
		return NewEmailMessage()
	case LOG_NOTIFY_TYPE:
		return NewLogMessage()
	default:
		return nil, errors.New(util.ENV_NOTIFY_TYPE + "must be defined")
	}
}

func (content NotifierContent) toString() string {
	return "Unit: " + content.Unit
}
