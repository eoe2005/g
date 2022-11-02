package gweb

import (
	"github.com/eoe2005/g/gconf"
	"github.com/gin-gonic/gin"
)

func getEncryptMiddleWare(conf *gconf.GWebEncryptYaml) gin.HandlerFunc {
	switch conf.Driver {
	case "aes":
		return getAesMiddleWare(conf)
	case "ssl":
		return getSslMiddleWare(conf)
	}
	return nil
}
func getAesMiddleWare(conf *gconf.GWebEncryptYaml) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.Next()
	}
}

func getSslMiddleWare(conf *gconf.GWebEncryptYaml) gin.HandlerFunc {
	return nil
}
