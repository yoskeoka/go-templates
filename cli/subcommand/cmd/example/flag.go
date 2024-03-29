package main

import (
	"fmt"
	"time"
)

type DateFlag struct {
	date *time.Time
}

func (d *DateFlag) Set(v string) error {
	t, err := time.Parse("20060102", v)
	if err == nil {
		*d.date = t
		return nil
	}

	return fmt.Errorf("-date format doesn't match any of supported format, where the supported format is 'yyyymmdd'")
}

func (d DateFlag) String() string {
	if d.date == nil {
		return "0000101"
	}
	return d.date.Format("20060102")
}
