package example

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerUser(rg *gin.Engine) {
	rg.Any("/ping", func(ctx *gin.Context) {
		ctx.Json(http.StatusOK, gin.H{
			"ping": "ok",
		})
	})
}
