package glog

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AccessLog() gin.HandlerFunc {
	ReOpenFile("access", webFd, func(fd *os.File) {
		webFd = fd
		if accessLog == nil {
			accessLog = log.New(fd, "HTTP:", log.Ldate|log.Ltime)
		} else {
			accessLog.SetOutput(fd)
		}
	})
	return func(ctx *gin.Context) {
		currentTime := time.Now()
		requestID := uuid.New()
		ctx.Set("request_id", requestID.String())
		ctx.Set("request_start_time", currentTime)
		ctx.Next()
		accessLog.Printf(" %s %d %d %s %s\n", ctx.Request.Method, ctx.Writer.Status(), time.Now().Sub(currentTime).Milliseconds(), ctx.ClientIP(), ctx.Request.RequestURI)
		ctx.Done()
	}
}
