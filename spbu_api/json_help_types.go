package spbu_api

import (
	"encoding/json"
	"strings"
	"time"
)

type JsonDate time.Time

func (j *JsonDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		return nil
	}
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = JsonDate(t)
	return nil
}

func (j JsonDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(j)
}

func (j JsonDate) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}

type ZLessTime time.Time

func (z *ZLessTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		return nil
	}
	t, err := time.Parse("2006-01-02T15:04:05", s)
	if err != nil {
		return err
	}
	*z = ZLessTime(t)
	return nil
}

func (z ZLessTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(z)
}

func (z ZLessTime) Format(s string) string {
	t := time.Time(z)
	return t.Format(s)
}
