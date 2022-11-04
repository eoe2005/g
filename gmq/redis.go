package gmq

import (
	"context"
	"fmt"
	"time"

	"github.com/eoe2005/g/gcache"
	"github.com/go-redis/redis/v8"
)

func getRedis(key string) *redis.Client {
	if r, ok := _localResis[key]; ok {
		return r
	}
	if c, ok := _localRedisConfMap[key]; ok {
		r := gcache.GetRedis(c.RefRedis)
		_localResis[key] = r
		return r
	}
	panic("没有mq的redis配置")
}
func buildRedisKey(name, topic string) string {
	return fmt.Sprintf("queue:%s:%s", name, topic)
}
func buildRedisAtKey(name, topic string) string {
	return fmt.Sprintf("queue:%s:%s:at", name, topic)
}
func buildRedisAfterKey(name, topic string) string {
	return fmt.Sprintf("queue:%s:%s:after", name, topic)
}
func RedisPublish(name, topic string, msg ...any) {
	getRedis(name).LPush(context.Background(), buildRedisKey(name, topic), msg...)
}
func RedisPublishAt(name, topic string, t time.Time, msg any) {
	getRedis(name).ZAdd(context.Background(), buildRedisAtKey(name, topic), &redis.Z{
		Score:  float64(t.UnixNano()),
		Member: msg,
	})
}
func RedisPublishAfter(name, topic string, t time.Duration, msg any) {
	getRedis(name).ZAdd(context.Background(), buildRedisAfterKey(name, topic), &redis.Z{
		Score:  float64(time.Now().Add(t).UnixNano()),
		Member: msg,
	})
}

func RedisConsumer(name, topic string, handle func(msg string) error) {
	redis := getRedis(name)
	c := context.Background()
	for {
		ss, e := redis.BRPop(c, time.Second*10, buildRedisKey(name, topic)).Result()
		if e != nil {
			panic(e.Error())
		}
		for _, v := range ss {
			handle(v)
		}
	}
}
func RedisConsumerAt(name, topic string, handle func(msg any) error) {
	r := getRedis(name)
	c := context.Background()
	rk := buildRedisAtKey(name, topic)
	for {
		ss, e := r.ZRangeByScoreWithScores(c, rk, &redis.ZRangeBy{
			Min:    "0",
			Max:    fmt.Sprintf("%d", time.Now().UnixNano()),
			Offset: 0, Count: 100,
		}).Result()
		if e != nil {
			panic(e.Error())
		}
		for _, v := range ss {

			if handle(v.Member) == nil {
				r.ZRem(c, rk, v.Member)
			}
		}
	}
}
func RedisConsumerAfter(name, topic string, handle func(msg any) error) {
	r := getRedis(name)
	c := context.Background()
	rk := buildRedisAfterKey(name, topic)
	for {
		ss, e := r.ZRangeByScoreWithScores(c, rk, &redis.ZRangeBy{
			Min:    "0",
			Max:    fmt.Sprintf("%d", time.Now().UnixNano()),
			Offset: 0, Count: 100,
		}).Result()
		if e != nil {
			panic(e.Error())
		}
		for _, v := range ss {

			if handle(v.Member) == nil {
				r.ZRem(c, rk, v.Member)
			}
		}
	}
}
