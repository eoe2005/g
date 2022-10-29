package g

import (
	"net/http"

	"github.com/eoe2005/g/gerror"
	"github.com/eoe2005/g/glog"
	"github.com/gin-gonic/gin"
)

func RunWeb(routerRegister func(*gin.Engine)) {
	r := gin.New()
	routerRegister(r)
	r.Use(
		glog.AccessLog(),
		func(ctx *gin.Context) {
			defer func() {
				err := recover()
				switch err.(type) {
				case gerror.JsonError:
					ctx.JSON(http.StatusOK, err)
				}
			}()
			ctx.Next()
		},
	)
	r.Run()
}
