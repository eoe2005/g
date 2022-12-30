package glog

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	applog = GetLog("app").NewLog("", log.LstdFlags)
	errlog = GetLog("error").NewLog("", log.LstdFlags)
)

func Error(ctx *gin.Context, format string, args ...any) {
	_writeLog(ctx, "ERROR", format, args...)
}
func Debug(ctx *gin.Context, format string, args ...any) {
	_writeLog(ctx, "DEBUG", format, args...)
}
func Info(ctx *gin.Context, format string, args ...any) {
	_writeLog(ctx, "INFO", format, args...)
}
func Waring(ctx *gin.Context, format string, args ...any) {
	_writeLog(ctx, "WARING", format, args...)
}

func _writeLog(ctx *gin.Context, level, format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	requireID := ""
	if ctx != nil {
		requireID = ctx.Value("request_id").(string)
	}
	applog.Printf("%s %s %s", level, requireID, msg)
}
