package gconf

import (
	"github.com/eoe2005/g/genv"
	"gopkg.in/yaml.v3"
)

var (
	confData = map[string][]byte{}
)

func GetAppConf() AppConf {
	return localConf
}
func RegisterConf(data map[string][]byte) {
	confData = data
}
func Load(target ...interface{}) {
	yamlContent, _ := confData[genv.GetRunEnv()]
	e := yaml.Unmarshal(yamlContent, &localConf)
	if e != nil {
		panic(e)
	}
	for _, item := range target {
		yaml.Unmarshal(yamlContent, item)
	}
}
