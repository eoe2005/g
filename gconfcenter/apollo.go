package gconfcenter

import (
	"strings"

	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
	"github.com/apolloconfig/agollo/v4/storage"
	"github.com/eoe2005/g/gconf"
)

func initApollo(c *gconf.GCfgYaml) {
	client, _ := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return &config.AppConfig{
			AppID:          c.AppID,
			Cluster:        c.Cluster,
			IP:             c.Host,
			NamespaceName:  c.Namespace,
			IsBackupConfig: false,
			Secret:         c.UserName,
		}, nil
	})

	client.AddChangeListener(apoloLister{Name: c.Name})
	ns := strings.Split(c.Namespace, ",")
	for _, n := range ns {
		client.GetConfigCache(n).Range(func(key, value interface{}) bool {
			_localCache.Set(c.Name+"."+n+"."+key.(string), value, 0)
			return true
		})
	}
}

type apoloLister struct {
	Name string
}

func (c apoloLister) OnChange(event *storage.ChangeEvent) {
	for k, i := range event.Changes {
		switch i.ChangeType {
		case storage.ADDED:
			_localCache.Set(c.Name+"."+event.Namespace+"."+k, i.NewValue, 0)
		case storage.DELETED:
			_localCache.Del(c.Name + "." + event.Namespace + "." + k)
		case storage.MODIFIED:
			_localCache.Set(c.Name+"."+event.Namespace+"."+k, i.NewValue, 0)
		}
	}
}
func (c apoloLister) OnNewestChange(event *storage.FullChangeEvent) {
	for k, i := range event.Changes {
		_localCache.Set(c.Name+"."+event.Namespace+"."+k, i, 0)
	}
}
