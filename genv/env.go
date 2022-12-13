package genv

import (
	"os"
	"path/filepath"
)

var (
	_appDir = ""
)

func GetAppDir() string {
	if _appDir == "" {
		filePath, _ := filepath.Abs(os.Args[0])
		_appDir = filepath.Dir(filePath)

	}
	return _appDir
}

func GetLogDir() string {
	if IsDebug() {
		dir, e := os.Getwd()
		if e != nil {
			return "/tmp/"
		}
		return dir + "/outer/logs/"
	}
	return GetAppDir() + "/../logs/"
}

func IsDebug() bool {
	logdir := GetAppDir() + "/../logs/"
	f, e := os.Open(logdir)
	if e != nil {
		return true
	}
	defer f.Close()
	return false
}
