package output

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JsonFail(g *gin.Context, code int, msg string) {
	g.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": gin.H{},
	})
	g.Done()
}
func JsonSuccess(g *gin.Context, data interface{}) {
	g.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
	g.Done()
}
