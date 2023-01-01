package glog

import (
	"context"
	"fmt"
	"log"
)

var (
	applog *log.Logger
)

func getapplog(name string) *log.Logger {
	if applog == nil {
		applog = GetLog(name).NewLog("", log.LstdFlags)
	}
	return applog
}
func Error(ctx context.Context, format string, args ...any) {
	_writeLog(ctx, "ERROR", format, args...)
}
func Debug(ctx context.Context, format string, args ...any) {
	_writeLog(ctx, "DEBUG", format, args...)
}
func Info(ctx context.Context, format string, args ...any) {
	_writeLog(ctx, "INFO", format, args...)
}
func Waring(ctx context.Context, format string, args ...any) {
	_writeLog(ctx, "WARING", format, args...)
}

func _writeLog(ctx context.Context, level, format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	requireID := ""
	if ctx != nil {
		requireID = ctx.Value("request_id").(string)
	}
	getapplog("app").Printf("%s %s %s", level, requireID, msg)
}
