package domain

import (
	"strings"
	"time"
)

type TrainTimestamp time.Time

func (ts *TrainTimestamp) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	t, err := time.ParseInLocation("2006-01-02T15:04:05", s, time.Local)

	if err != nil {
		return err
	}
	*ts = TrainTimestamp(t)
	return nil
}

func (ts TrainTimestamp) String() string {
	return time.Time(ts).Format(time.Kitchen)
}

type BusTimestamp time.Time

func (ts *BusTimestamp) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	t, err := time.ParseInLocation("20060102 15:04", s, time.Local)

	if err != nil {
		return err
	}
	*ts = BusTimestamp(t)
	return nil
}

func (ts BusTimestamp) String() string {
	return time.Time(ts).Format(time.Kitchen)
}
