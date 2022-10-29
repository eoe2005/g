package gcache

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/eoe2005/g/gconf"
	"github.com/go-redis/redis/v8"
)

var (
	_localRedis        = map[string]*redis.Client{}
	_localRedisCluster = map[string]*redis.ClusterClient{}
	_localMemcached    = map[string]*memcache.Client{}
)

func Register(cacheList []*gconf.GConfYaml) {
	for _, item := range cacheList {
		switch item.Driver {
		case "redis":
			_localRedis[item.Name] = conRedis(item)
		case "redisCluster":
			_localRedisCluster[item.Name] = conRedisCluster(item)
		case "memcache":
			_localMemcached[item.Name] = conMemcache(item)
		}
	}
}

func GetRedis(key string) *redis.Client {
	if r, ok := _localRedis[key]; ok {
		return r
	}
	panic("没有配置")
}
func GetRedisCluster(key string) *redis.ClusterClient {
	if r, ok := _localRedisCluster[key]; ok {
		return r
	}
	panic("没有配置")
}
func GetMemcache(key string) *memcache.Client {
	if r, ok := _localMemcached[key]; ok {
		return r
	}
	panic("没有配置")
}
