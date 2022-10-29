package gconfcenter

import (
	"strconv"

	"github.com/eoe2005/g/gconf"
)

var (
	_localCache = &confCache{}
)

func Register(cs []*gconf.GCfgYaml) {
	for _, c := range cs {
		switch c.Driver {
		case "apollo":
			initApollo(c)
		}
	}
}

func Get(key string) interface{} {
	r, e := _localCache.Get(key)
	if e != nil {
		return nil
	}
	return r
}
func GetString(key, defval string) string {
	r := Get(key)
	if r == nil {
		return defval
	}
	return r.(string)
}
func GetInt(key string, defval int64) int64 {
	r := GetString(key, "")
	if r == "" {
		return defval
	}
	r1, e := strconv.ParseInt(r, 10, 64)
	if e != nil {
		return defval
	}
	return r1
}
func GetFloat(key string, defval float64) float64 {
	r := GetString(key, "")
	if r == "" {
		return defval
	}
	r1, e := strconv.ParseFloat(r, 10)
	if e != nil {
		return defval
	}
	return r1
}
