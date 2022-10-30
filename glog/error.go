package glog

import (
	"os"
)

func RegisterErrorLog() {
	ReOpenFile("error", errFd, func(fd *os.File) {
		os.Stderr = fd
		os.Stdout = fd
	})
}
