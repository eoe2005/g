package gmq

import (
	"context"
	"encoding/json"
	"errors"
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

func PulishKfakaMsg(ctx context.Context, name string, msg ...interface{}) error {
	if len(msg) == 0 {
		return nil
	}
	if k, ok := _localKafkaWriter[name]; ok {
		sendMsg := []kafka.Message{}
		for _, m := range msg {
			val := []byte("")
			switch m.(type) {
			case string:
				val = []byte(m.(string))
			case []byte:
				val = m.([]byte)
			default:
				val, _ = json.Marshal(m)
			}
			sendMsg = append(sendMsg, kafka.Message{
				Value: val,
			})
		}
		return k.WriteMessages(ctx, sendMsg...)
	}
	return errors.New("配置不存在")
}
