package output

import (
	"net/http"

	"github.com/eoe2005/g/gtemplate"
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

// 显示json
func JsonSuccess(g *gin.Context, data any) {
	g.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
	g.Done()
}

//展示html
func Display(g *gin.Context, tpl string, data any) {
	g.Status(http.StatusOK)
	g.Writer.WriteString(gtemplate.Fetch(tpl, data))
	g.Done()
}
