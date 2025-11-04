package utils

import (
	"taska-core-me-go/cmd/api/constants"
	"time"
)

func FormatDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(constants.DateFormatISO)
}

func FormatDateTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(constants.DateTimeFormatFull)
}
