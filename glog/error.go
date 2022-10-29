package glog

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Register(g *gin.Engine) {
	ReOpenFile("error", errFd, func(fd *os.File) {
		os.Stderr = fd
		os.Stdout = fd
	})
}
