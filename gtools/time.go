package gtools

import "time"

func UnixToDatetimeString(t int64) string {
	return time.Unix(t, 0).Format("2006-01-02 15:04:05")
}
func UnixMToDatetimeString(t int64) string {
	return time.UnixMilli(t).Format("2006-01-02 15:04:05")
}
