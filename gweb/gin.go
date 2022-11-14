package gweb

import (
	"fmt"

	"github.com/eoe2005/g/gconf"
	"github.com/gin-gonic/gin"
)

var (
	_ginMiddleWare = map[string]gin.HandlerFunc{}
)

func GetMiddleWareByName(name ...string) []gin.HandlerFunc {
	ret := []gin.HandlerFunc{}
	for _, n := range name {
		if h, ok := _ginMiddleWare[n]; ok {
			ret = append(ret, h)
		}
	}
	return ret
}
func Register(confs *gconf.GWebYaml) {
	fmt.Sprintf("conf %v \n", confs)
	if confs == nil {
		return
	}
	for _, conf := range confs.MiddleWare {
		h := initMiddleWare(conf)
		if h != nil {
			_ginMiddleWare[conf.Name] = h
		}
	}
}
func initMiddleWare(conf *gconf.GWebMiddleWareYaml) gin.HandlerFunc {
	switch conf.Driver {
	case "jwt":
		return initJwtMiddleWare(conf)
	case "session_redis":
		return initRedisMiddleWare(conf)
	case "session_redis_cluster":
		return initRedisClusterMiddleWare(conf)
	case "aes":
		return initAesMiddleWare(conf)
	case "rsa":
		return initRsaMiddleWare(conf)
	}
	return nil
}
