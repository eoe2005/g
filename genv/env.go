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

func GetAppConfFile(name string) string {
	return GetAppDir() + "../conf/" + name + ".yaml"
}

func GetLogDir() string {
	return GetAppDir() + "../logs/"
}
