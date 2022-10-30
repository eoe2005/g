package g

import (
	"github.com/eoe2005/g/glog"
	"github.com/gin-gonic/gin"
)

func RunWeb(routerRegister func(*gin.Engine)) {
	// glog.RegisterErrorLog()
	r := gin.New()
	routerRegister(r)
	r.Use(
		glog.AccessLog(),
		// func(ctx *gin.Context) {
		// 	defer func() {
		// 		err := recover()
		// 		switch err.(type) {
		// 		case gerror.JsonError:
		// 			ctx.JSON(http.StatusOK, err)
		// 		}
		// 	}()
		// 	ctx.Next()
		// },
		// gin.Recovery(),
	)
	r.Run("0.0.0.0:8888")
}
