package gcache

import (
	"strings"

	"github.com/eoe2005/g/gconf"
	"github.com/go-redis/redis/v8"
)

func conRedis(c *gconf.GConfYaml) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.Host,
		Username: c.UserName,
		Password: c.UserPass,
		DB:       c.DB,
	})
}
func conRedisCluster(c *gconf.GConfYaml) *redis.ClusterClient {
	return redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    strings.Split(c.Host, ","),
		Username: c.UserName,
		Password: c.UserPass,
	})
}
