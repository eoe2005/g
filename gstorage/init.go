package gstorage

import "github.com/eoe2005/g/gconf"

var (
	_localStorageConf = map[string]*gconf.GStorageYaml{}
)

func Register(sList []*gconf.GStorageYaml) {
	for _, item := range sList {
		_localStorageConf[item.Name] = item
	}
}
