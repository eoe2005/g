package g

import (
	"github.com/eoe2005/g/glog"
	"github.com/gin-gonic/gin"
)

func runWeb() {
	r := gin.New()

	r.Use(
		glog.AccessLog(),
	)
}
