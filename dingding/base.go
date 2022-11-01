package dingding

import "github.com/eoe2005/g/gconf"

const ddapihost = "https://oapi.dingtalk.com/robot/send?access_token="

var (
	_localConf = map[string]*gconf.GDingDingYaml{}
)

func Register(confs []*gconf.GDingDingYaml) {
	for _, i := range confs {
		_localConf[i.Name] = i
	}
}

func send(name string, msg any) {
	url := ""
	if c, ok := _localConf[name]; ok {
		url = ddapihost + c.Token
	}
	if url == "" {
		return
	}

}
