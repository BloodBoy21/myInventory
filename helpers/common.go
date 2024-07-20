package helpers

import "reflect"

func MergeStruct(source, destination interface{}) {
	dst := reflect.ValueOf(destination).Elem()
	src := reflect.ValueOf(source).Elem()
	for i := 0; i < src.NumField(); i++ {
		if !reflect.DeepEqual(src.Field(i).Interface(), reflect.Zero(src.Field(i).Type()).Interface()) {
			dst.Field(i).Set(src.Field(i))
		}
	}
}
