package glog

import (
	"os"
	"strings"

	"github.com/eoe2005/g/gconf"
)

func GetLog(fname string) *Glog {
	localConf := gconf.GetAppConf()
	if localConf.Log != nil {
		fileName := localConf.Log.Dir + "/" + fname + ".log"
		switch strings.ToLower(localConf.Log.SplitType) {
		case "filesize":
			return NewGlogFileSize(fileName, localConf.Log.MaxFileSize)
		case "hour":
			return NewGlogSplitHour(fileName)
		default:
			return NewGlogSplitDay(fileName)
		}
	}
	dir, e := os.Getwd()
	if e != nil {
		dir = "/tmp"
	}
	return NewGlogSplitDay(dir + "/" + fname + ".log")
}
