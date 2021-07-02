package util

import (
	"fmt"
	"time"
)

// returns an ISO8601 formatted string
// YYYYMMDDTHHMMSSZ
func FormatTimeStamp(t time.Time) string {
	return fmt.Sprintf("%d%02d%02dT%02d%02d%02dZ", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}
