package utils

import (
	"fmt"
	"strings"
	"time"
)

const TimeLayout = "2006-01-02T15:04:05-07:00"

type JSONTime struct {
	time.Time
}

func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", t.Format(TimeLayout))
	return []byte(stamp), nil
}

func ParseTime(date string) (time.Time, error) {
	// URL decoding of a '+' is a whitespace, so we change it back to '+'
	date = strings.ReplaceAll(date, " ", "+")
	return time.Parse(TimeLayout, date)
}
