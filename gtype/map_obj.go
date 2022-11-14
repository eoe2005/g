package gtype

import (
	"errors"
	"reflect"
)

var (
	_km = []string{"db", "json", "cache", "field", "form", "get", "post"}
)

// map转对象
func MapStringToObj(src map[string]string, target any) error {
	t := reflect.ValueOf(target)
	if t.Kind() != reflect.Ptr {
		return errors.New("类型错误")
	}
	t = t.Elem()
	if t.Kind() != reflect.Struct {
		return errors.New("类型错误")
	}
	flen := t.NumField()

	for i := 0; i < flen; i++ {
		v := t.Field(i)
		tg := t.Type().Field(i).Tag
		for _, k := range _km {
			tn := tg.Get(k)
			if tn == "" {
				continue
			}
			if srv, ok := src[tn]; ok {
				SetStringToValue(srv, v)
				break
			}
		}
	}
	return nil
}
