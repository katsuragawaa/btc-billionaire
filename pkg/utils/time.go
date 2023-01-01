package utils

import (
	"fmt"
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
