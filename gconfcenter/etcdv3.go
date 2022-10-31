package gconfcenter

import (
	"context"
	"strings"
	"time"

	"github.com/eoe2005/g/gconf"
	etcdv3 "go.etcd.io/etcd/client/v3"
)

func initEtcd3(c *gconf.GCfgYaml) {
	client, err := etcdv3.New(etcdv3.Config{
		Endpoints:   strings.Split(c.Host, ","),
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return
	}
	kv := etcdv3.NewKV(client)
	res, e := kv.Get(context.Background(), "/"+c.Namespace, etcdv3.WithPrefix())
	if e != nil {
		return
	}
	for _, i := range res.Kvs {
		_localCache.Set(c.Name+"."+string(i.Key), string(i.Value), 0)
	}
	go watchEtcd3(c, client)

}
func watchEtcd3(c *gconf.GCfgYaml, client *etcdv3.Client) {
	wc := client.Watch(context.TODO(), "/"+c.Namespace, etcdv3.WithPrefix())
	for v := range wc {
		for _, e := range v.Events {
			if e.Type == etcdv3.EventTypePut {
				_localCache.Del(c.Name + "." + string(e.Kv.Key))
			} else {
				_localCache.Set(c.Name+"."+string(e.Kv.Key), string(e.Kv.Value), 0)
			}
		}
	}

}
