package glog

import (
	"os"
	"strings"

	"github.com/eoe2005/g/gconf"
)

var (
	_logMap = map[string]*Glog{}
)

func GetLog(fname string) *Glog {
	r, ok := _logMap[fname]
	if ok {
		return r
	}
	localConf := gconf.GetAppConf()
	if localConf.Log != nil {
		fileName := localConf.Log.Dir + "/" + fname + ".log"
		switch strings.ToLower(localConf.Log.SplitType) {
		case "filesize":
			r = NewGlogFileSize(fileName, localConf.Log.MaxFileSize)
			_logMap[fname] = r
			return r
		case "hour":
			r = NewGlogSplitHour(fileName)
			_logMap[fname] = r
			return r
		default:
			r = NewGlogSplitDay(fileName)
			_logMap[fname] = r
			return r

		}
	}
	dir, e := os.Getwd()
	if e != nil {
		dir = "/tmp"
	}
	r = NewGlogSplitDay(dir + "/" + fname + ".log")
	_logMap[fname] = r
	return r
}
