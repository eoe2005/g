package gweb

import (
	"github.com/eoe2005/g/gconf"
	"github.com/gin-gonic/gin"
)

var (
	mids = []gin.HandlerFunc{}
)

func Register(conf *gconf.GWebYaml) {
	if conf.Auth != nil {
		h := getAuthMiddleWare(conf.Auth)
		if h != nil {
			mids = append(mids, h)
		}
	}
	if conf.Encrypt != nil {
		h := getEncryptMiddleWare(conf.Encrypt)
		if h != nil {
			mids = append(mids, h)
		}
	}
}
func GetMiddleWare() []gin.HandlerFunc {
	return mids
}
