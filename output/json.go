package output

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JsonFail(g *gin.Context, code int, msg string) {
	json(g, code, msg, gin.H{})
}
func JsonSuccess(g *gin.Context, data interface{}) {
	json(g, 0, "success", data)
}

func json(g *gin.Context, code int, msg string, data interface{}) {
	g.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}
