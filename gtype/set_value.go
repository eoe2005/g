package gtype

import (
	"reflect"
	"strconv"
	"time"
)

func SetStringToValue(src string, val reflect.Value) {
	switch val.Type().Kind() {
	case reflect.String:
		val.SetString(src)
	case reflect.Float32, reflect.Float64:
		f, _ := strconv.ParseFloat(src, 64)
		val.SetFloat(f)
	case reflect.Bool:
		f, _ := strconv.ParseBool(src)
		val.SetBool(f)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		f, _ := strconv.ParseInt(src, 10, 64)
		val.SetInt(f)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		f, _ := strconv.ParseUint(src, 10, 64)
		val.SetUint(f)
	default:
		if val.Type().Name() == "time.Time" {
			t, _ := time.Parse("2006-01-02 15:04:05", src)
			val.Set(reflect.ValueOf(t))
		}
	}
}
