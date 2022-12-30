package gcache

import (
	"strings"

	"github.com/eoe2005/g/gconf"
	"github.com/eoe2005/g/gmonitor"
	"github.com/go-redis/redis/v8"
)

func conRedis(c *gconf.GCacheYaml) *redis.Client {
	ret := redis.NewClient(&redis.Options{
		Addr:     c.Host,
		Username: c.UserName,
		Password: c.UserPass,
		DB:       c.DB,
	})
	ret.AddHook(gmonitor.NewRedisPlugin(c.Host))
	return ret
}
func conRedisCluster(c *gconf.GCacheYaml) *redis.ClusterClient {
	ret := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    strings.Split(c.Host, ","),
		Username: c.UserName,
		Password: c.UserPass,
	})
	ret.AddHook(gmonitor.NewRedisPlugin(c.Host))
	return ret
}
