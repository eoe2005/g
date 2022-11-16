package gmonitor

import (
	"context"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

func NewRedisPlugin(host string) redis.Hook {
	return &redisMetrics{Host: host}
}

type redisMetrics struct {
	Host string
}

func (p *redisMetrics) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	ctx = context.WithValue(ctx, "startTime", time.Now())
	return ctx, nil
}
func (p *redisMetrics) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	s := ctx.Value("startTime")
	st, e := s.(time.Time)
	if !e {
		return nil
	}
	Summary("redis_cost", "redis_cost", float64(time.Now().Sub(st)/time.Millisecond), map[string]string{
		"command":   cmd.Name(),
		"server_ip": p.Host,
	})
	return nil
}

func (p *redisMetrics) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	ctx = context.WithValue(ctx, "startTime", time.Now())
	return ctx, nil
}
func (p *redisMetrics) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
	s := ctx.Value("startTime")
	st, e := s.(time.Time)
	if !e {
		return nil
	}
	command := []string{}
	for _, cc := range cmds {
		command = append(command, cc.Name())
	}
	Summary("redis_cost", "redis_cost", float64(time.Now().Sub(st)/time.Millisecond), map[string]string{
		"command":   strings.Join(command, ","),
		"server_ip": p.Host,
	})
	return nil
}
