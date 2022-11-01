package gmq

import (
	"github.com/eoe2005/g/gconf"
	"github.com/segmentio/kafka-go"
)

var (
	_localKafkaWriter = map[string]*kafka.Writer{}
)

func Register(mqs []*gconf.GMqYaml) {
	for _, c := range mqs {
		switch c.Driver {
		case "kafka":
			_localKafkaWriter[c.Name] = conKafkaWrite(c)
		case "rabbitmq":
		}
	}
}
