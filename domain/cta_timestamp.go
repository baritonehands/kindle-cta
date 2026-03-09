package domain

import (
	"strings"
	"time"
)

type CtaTimestamp time.Time

func (ts *CtaTimestamp) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	t, err := time.ParseInLocation("2006-01-02T15:04:05", s, time.Local)

	if err != nil {
		return err
	}
	*ts = CtaTimestamp(t)
	return nil
}

func (ts CtaTimestamp) String() string {
	return time.Time(ts).Format(time.Kitchen)
}
