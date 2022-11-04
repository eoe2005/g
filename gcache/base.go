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

	_localConfMap = map[string]*gconf.GCacheYaml{}
)

func Register(cacheList []*gconf.GCacheYaml) {
	for _, item := range cacheList {
		_localConfMap[item.Name] = item
	}
}

func GetRedis(key string) *redis.Client {
	if r, ok := _localRedis[key]; ok {
		return r
	}
	if conf, ok := _localConfMap[key]; ok {
		if conf.Driver == "redis" {
			r := conRedis(conf)
			_localRedis[key] = r
			return r
		}
	}
	panic("没有配置")
}
func GetRedisCluster(key string) *redis.ClusterClient {
	if r, ok := _localRedisCluster[key]; ok {
		return r
	}
	if conf, ok := _localConfMap[key]; ok {
		if conf.Driver == "redisCluster" {
			r := conRedisCluster(conf)
			_localRedisCluster[key] = r
			return r
		}
	}
	panic("没有配置")
}
func GetMemcache(key string) *memcache.Client {
	if r, ok := _localMemcached[key]; ok {
		return r
	}
	if conf, ok := _localConfMap[key]; ok {
		if conf.Driver == "memcache" {
			r := conMemcache(conf)
			_localMemcached[key] = r
			return r
		}
	}
	panic("没有配置")
}
