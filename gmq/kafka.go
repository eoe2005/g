package gmq

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/eoe2005/g/gconf"
	"github.com/eoe2005/g/glog"
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
	var kc *kafka.Writer
	if k, ok := _localKafkaWriter[name]; ok {
		kc = k
	} else {
		if conf, ok := _localKafkaConfMap[name]; ok {
			kc = conKafkaWrite(conf)
			_localKafkaWriter[name] = kc
		} else {
			return errors.New("配置不存在")
		}
	}
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
	return kc.WriteMessages(ctx, sendMsg...)

}

func KafkaConsumer(ctx context.Context, groupname, name string, handle func(msg kafka.Message) error, isAutoComit bool) {
	conf, ok := _localKafkaConfMap[name]
	if !ok {
		glog.Error(ctx, "kafka consumer %s conf is not", name)
		return
	}
	kConf := kafka.ReaderConfig{
		Brokers:               strings.Split(conf.Hosts, ","),
		GroupID:               groupname,
		GroupTopics:           []string{conf.Topic},
		MinBytes:              10e3, // 10KB
		MaxBytes:              10e6, // 10MB
		SessionTimeout:        time.Second * 60,
		WatchPartitionChanges: true,
		// HeartbeatInterval: time.Second,
		MaxWait:     time.Second * 5,
		StartOffset: kafka.LastOffset, //勇飞建议设置
		// CommitInterval:    time.Second,
		// Logger:      logs.GetBusinessLogger(),
		// ErrorLogger: logs.GetBusinessLogger(),
		// MaxAttempts:       10,
		Dialer: &kafka.Dialer{
			Timeout:   time.Hour * 5,
			DualStack: true,
			KeepAlive: time.Hour * 5,
		},
	}
	r := kafka.NewReader(kConf)
	defer r.Close()
	for {
		if r.Stats().Errors > 0 {
			r.Close()
			r = kafka.NewReader(kConf)
		}
		if isAutoComit {
			msg, e := r.ReadMessage(ctx)
			if e != nil {
				break
			}
			go func() {
				defer func() {
					if e := recover(); e != nil {
						glog.Error(ctx, "kafka [%s] error %v", groupname, e)
					}
				}()
				handle(msg)
			}()
		} else {
			msg, e := r.FetchMessage(ctx)
			if e != nil {
				break
			}
			func() {
				defer func() {
					if e := recover(); e != nil {
						glog.Error(ctx, "kafka [%s] error %v", groupname, e)
					} else {
						r.CommitMessages(ctx, msg)
					}
				}()
				e := handle(msg)
				if e != nil {
					panic(e)
				}
			}()
		}
	}
}
