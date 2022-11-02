package g

import (
	"github.com/eoe2005/g/glog"
	"github.com/eoe2005/g/gweb"
	"github.com/gin-gonic/gin"
)

func RunWeb(routerRegister func(*gin.Engine)) {
	initConfig()
	// glog.RegisterErrorLog()
	r := gin.New()
	mids := []gin.HandlerFunc{
		glog.AccessLog(),
	}
	mids = append(mids, gweb.GetMiddleWare()...)
	mids = append(mids, gin.Recovery())
	r.Use(mids...)
	routerRegister(r)
	r.Run("0.0.0.0:8888")
}
