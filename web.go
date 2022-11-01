package g

import (
	"github.com/eoe2005/g/glog"
	"github.com/gin-gonic/gin"
)

func RunWeb(routerRegister func(*gin.Engine)) {
	initConfig()
	// glog.RegisterErrorLog()
	r := gin.New()

	r.Use(
		glog.AccessLog(),

		gin.Recovery(),
	)
	routerRegister(r)
	r.Run("0.0.0.0:8888")
}
