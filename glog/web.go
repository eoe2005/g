package glog

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AccessLog() gin.HandlerFunc {
	logf := GetLog("access").NewLog("", log.LstdFlags)
	return func(ctx *gin.Context) {
		currentTime := time.Now()
		requestID := uuid.New()
		ctx.Set("request_id", requestID.String())
		ctx.Set("request_start_time", currentTime)
		ctx.Next()
		ctx.Header("request_id", requestID.String())
		logf.Printf(" %s %d %d %s %s", ctx.Request.Method, ctx.Writer.Status(), time.Now().Sub(currentTime).Milliseconds(), ctx.ClientIP(), ctx.Request.RequestURI)
		// saveLog("access", fmt.Sprintf(" %s %d %d %s %s", ctx.Request.Method, ctx.Writer.Status(), time.Now().Sub(currentTime).Milliseconds(), ctx.ClientIP(), ctx.Request.RequestURI))
		// accessLog.Printf(" %s %d %d %s %s\n", ctx.Request.Method, ctx.Writer.Status(), time.Now().Sub(currentTime).Milliseconds(), ctx.ClientIP(), ctx.Request.RequestURI)

	}
}
