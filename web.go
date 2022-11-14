package g

import (
	"github.com/eoe2005/g/glog"
	"github.com/gin-gonic/gin"
)

func RunWeb(routerRegister func(*gin.Engine)) {
	initConfig()
	// glog.RegisterErrorLog()
	r := gin.New()
	mids := []gin.HandlerFunc{
		glog.AccessLog(),
		gin.Recovery(),
	}
	r.Use(mids...)
	for _, c := range localCall {
		c()
	}
	routerRegister(r)

	r.Run("0.0.0.0:8888")
}
