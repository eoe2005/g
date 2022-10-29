package g

import (
	"github.com/eoe2005/g/glog"
	"github.com/gin-gonic/gin"
)

func runWeb(routerRegister func(*gin.Engine)) {
	r := gin.New()
	routerRegister(r)
	r.Use(
		glog.AccessLog(),
	)
	r.Run()

}
