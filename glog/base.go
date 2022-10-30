package glog

import (
	"fmt"
	"log"
	"os"

	"github.com/eoe2005/g/genv"
)

var (
	errFd     *os.File
	logFd     *os.File
	webFd     *os.File
	accessLog *log.Logger
	localLog  *log.Logger
)

func ReOpenFile(name string, oldFd *os.File, callHandel func(fd *os.File)) error {
	filePath := genv.GetLogDir() + name + ".log"
	// os.Rename(filePath, genv.GetLogDir()+name + time.+".log")
	nf, e := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
	if e != nil {
		fmt.Printf("打开文件失败 %s\n", e.Error())
		return e
	}
	if oldFd != nf && oldFd != nil {
		oldFd.Close()
	}
	callHandel(nf)
	return nil
}
