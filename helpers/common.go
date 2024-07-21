package helpers

import (
	"reflect"
	"time"
)

func MergeStruct(source, destination interface{}) {
	dst := reflect.ValueOf(destination).Elem()
	src := reflect.ValueOf(source).Elem()
	for i := 0; i < src.NumField(); i++ {
		if !reflect.DeepEqual(src.Field(i).Interface(), reflect.Zero(src.Field(i).Type()).Interface()) {
			dst.Field(i).Set(src.Field(i))
		}
	}
}

func ParseDate(dateString string, location *time.Location) time.Time {
	date, err := time.ParseInLocation("2006-01-02", dateString, location)
	if err != nil {
		return time.Time{}
	}
	return date
}

func SetStartOfDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
}

func SetEndOfDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 0, date.Location())
}
