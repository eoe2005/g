package gstorage

import (
	"time"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func QiniuGetToken(confName string) string {
	conf, ok := _localStorageConf[confName]
	if !ok {
		return ""
	}
	if conf.Driver != "qiniu" {
		return ""
	}
	mc := qbox.NewMac(conf.Key, conf.Secret)
	p := storage.PutPolicy{
		Scope:      conf.Bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}
	return p.UploadToken(mc)
}
func QuniuPublicUrl(confName, file string) string {
	conf, ok := _localStorageConf[confName]
	if !ok {
		return ""
	}
	if conf.Driver != "qiniu" {
		return ""
	}
	return storage.MakePublicURLv2(conf.Domain, file)
}
func QuniuPrivateUrl(confName, file string) string {
	conf, ok := _localStorageConf[confName]
	if !ok {
		return ""
	}
	if conf.Driver != "qiniu" {
		return ""
	}
	mc := qbox.NewMac(conf.Key, conf.Secret)
	return storage.MakePrivateURLv2(mc, conf.Domain, file, time.Now().Add(time.Second*3600).Unix())
}
