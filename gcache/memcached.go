package gcache

import (
	"strings"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/eoe2005/g/gconf"
)

func conMemcache(c *gconf.GCacheYaml) *memcache.Client {
	return memcache.New(strings.Split(c.Host, ",")...)
}
