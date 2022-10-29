package gconfcenter

import (
	"time"

	etcdv3 "go.etcd.io/etcd/client/v3"
)

func conEtcd3() {
	_, err := etcdv3.New(etcdv3.Config{
		Endpoints:   []string{},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return
	}

}
