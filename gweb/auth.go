package gweb

import (
	"github.com/eoe2005/g/gconf"
	"github.com/gin-gonic/gin"
)

func getAuthMiddleWare(conf *gconf.GWebAuthYaml) gin.HandlerFunc {
	switch conf.Driver {
	case "jwt":
		return getJwtMiddleWare(conf)
	case "redis":
		return getRedisMiddleWare(conf)
	}
	return nil
}
func getJwtMiddleWare(conf *gconf.GWebAuthYaml) gin.HandlerFunc {
	return nil
}

func getRedisMiddleWare(conf *gconf.GWebAuthYaml) gin.HandlerFunc {
	return nil
}
