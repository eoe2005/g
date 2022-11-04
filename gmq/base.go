package gmq

import (
	"github.com/eoe2005/g/gconf"
	"github.com/go-redis/redis/v8"
	"github.com/segmentio/kafka-go"
)

var (
	_localKafkaWriter  = map[string]*kafka.Writer{}
	_localKafkaConfMap = map[string]*gconf.GMqYaml{}
	_localResis        = map[string]*redis.Client{}
	_localRedisConfMap = map[string]*gconf.GMqYaml{}
)

func Register(mqs []*gconf.GMqYaml) {
	for _, c := range mqs {
		switch c.Driver {
		case "kafka":
			_localKafkaConfMap[c.Name] = c
		case "redis":
			_localRedisConfMap[c.Name] = c
		}
	}
}
