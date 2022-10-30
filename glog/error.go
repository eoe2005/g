package glog

import (
	"fmt"
	"os"
)

func RegisterErrorLog() {
	ReOpenFile("error", errFd, func(fd *os.File) {
		fmt.Printf("初始化错误日志 %v \n", fd)
		// os.Stderr = fd
		// os.Stdout = fd
		fmt.Printf("初始化错误日志\n")
	})
}
