package glog

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
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

func saveLog(filename, data string) {
	fmt.Printf("%s %s\n", time.Now().Format(time.RFC3339Nano), data)
	// filePath := genv.GetLogDir() + filename + ".log"
	// data = fmt.Sprintf("%s %s\n", time.Now().Format(time.RFC3339Nano), data)
	// ioutil.WriteFile()
	// ioutil.WriteFile(filePath, []byte(data), fs.ModePerm|fs.ModeAppend)
}
func _writeLog(ctx *gin.Context, level, format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	requireID := ""
	if ctx != nil {
		requireID = ctx.Value("request_id").(string)
	}
	// if localLog == nil {
	// 	ReOpenFile("log", logFd, func(fd *os.File) {
	// 		localLog = log.New(fd, "", log.Ldate|log.Ltime)
	// 	})
	// }
	saveLog("app", fmt.Sprintf("%s %s %s", level, requireID, msg))
}
