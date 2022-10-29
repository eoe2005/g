package gmq

import (
	"strings"

	"github.com/eoe2005/g/gconf"
	"github.com/segmentio/kafka-go"
)

func conKafkaWrite(c *gconf.GMqYaml) *kafka.Writer {
	return &kafka.Writer{
		Addr:  kafka.TCP(strings.Split(c.Hosts, ",")...),
		Topic: c.Topic,
	}
}
