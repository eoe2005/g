package output

import (
	"github.com/eoe2005/g/gerror"
	"github.com/gin-gonic/gin"
)

func JsonFail(g *gin.Context, code int, msg string) {
	panic(gerror.JsonError{
		Code: code,
		Msg:  msg,
		Data: gin.H{},
	})
}
func JsonSuccess(g *gin.Context, data interface{}) {
	panic(gerror.JsonError{
		Code: 0,
		Msg:  "success",
		Data: gin.H{},
	})
}
