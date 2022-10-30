package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerRouter(rg *gin.Engine) {
	rg.Any("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"ping": "ok",
		})
	})
}
