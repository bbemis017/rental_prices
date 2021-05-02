package notifications

import (
	"errors"

	"github.com/bbemis017/ApartmentNotifier/util"
)

type (
	Notifier interface {
		Send()
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
		return NewLogMessage("testing message")
	default:
		return nil, errors.New(util.ENV_NOTIFY_TYPE + "must be defined")
	}
}
