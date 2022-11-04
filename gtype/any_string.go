package gtype

import "encoding/json"

func AnyToString(src any) string {
	b, e := json.Marshal(src)
	if e != nil {
		return ""
	}
	return string(b)
}
func StringToAny(src string, target any) error {
	return json.Unmarshal([]byte(src), target)
}
func ArrAnyToString(src []any) []string {
	ret := []string{}
	for _, i := range src {
		ret = append(ret, AnyToString(i))
	}
	return ret
}

// func ArrStringToAny(src []string, target any) []any {
// 	ret := []any{}
// 	for _, i := range src {
// 		t := target.Copy()
// 		StringToAny()
// 		ret = append(ret, (i))
// 	}
// 	return ret
// }
