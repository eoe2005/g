package gconf

import (
	"io/ioutil"
	"os"

	"github.com/eoe2005/g/genv"
	"gopkg.in/yaml.v3"
)

func GetAppConf() AppConf {
	return localConf
}
func Load(target ...interface{}) {
	file := genv.GetAppConfFile("conf-" + genv.GetRunEnv())
	fd, e := os.Open(file)
	if e != nil {
		panic("配置不存在")
	}
	defer fd.Close()
	yamlContent, e := ioutil.ReadAll(fd)
	if e != nil {
		panic("读取配置错误")
	}
	e = yaml.Unmarshal(yamlContent, &localConf)
	if e != nil {
		panic(e)
	}
	for _, item := range target {
		yaml.Unmarshal(yamlContent, item)
	}
}
