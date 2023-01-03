package gtools

import "strings"

func SplitUpperFirst(src string, splitName string) string {
	ret := []string{}
	sd := strings.Split(src, splitName)
	for _, val := range sd {
		val = strings.TrimSpace(val)
		if val == "" {
			continue
		}
		ret = append(ret, strings.ToUpper(val[0:1])+val[1:])
	}
	return strings.Join(ret, "")
}
func StrUFirst(src string) string {
	return strings.ToUpper(src[0:1]) + src[1:]
}
