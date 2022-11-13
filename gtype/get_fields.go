package gtype

import "reflect"

func FetchFeildsByName(target any, fname ...string) []map[string]reflect.Value {
	ret := []map[string]reflect.Value{}

	vv := reflect.ValueOf(target)
	for vv.Type().Kind() == reflect.Ptr {
		vv = vv.Elem()
	}
	switch vv.Type().Kind() {
	case reflect.Struct:
		ret = append(ret, getFieldByStruct(vv, fname...))
	case reflect.Array, reflect.Slice, reflect.Map:
		for i := 0; i < vv.Len(); i++ {
			ret = append(ret, getFieldByStruct(vv.Index(i), fname...))
		}
	}
	return ret
}

func getFieldByStruct(vv reflect.Value, fname ...string) map[string]reflect.Value {
	ret := map[string]reflect.Value{}

	for i := 0; i < vv.NumField(); i++ {
		fg := vv.Type().Field(i).Tag.Get("json")
		for _, name := range fname {
			if fg == name {
				ret[name] = vv.Field(i)
			} else {
				ret[name] = reflect.ValueOf(nil)
			}
		}
	}
	return ret
}
