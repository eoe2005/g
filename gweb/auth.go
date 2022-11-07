package gweb

import (
	"encoding/json"
	"net/url"
	"time"

	"github.com/eoe2005/g/gcache"
	"github.com/eoe2005/g/gconf"
	"github.com/eoe2005/g/glog"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wumansgy/goEncrypt/aes"
)

func initJwtMiddleWare(conf *gconf.GWebMiddleWareYaml) gin.HandlerFunc {
	return buildAuthMiddleWare(conf, func(s *string, sess *GSession, gay *gconf.GWebMiddleWareYaml, ctx *gin.Context) {
		if *s != "" {
			data, err := aes.AesCbcDecryptByBase64(*s, []byte(conf.AuthKey), nil)
			if err == nil {
				ee := json.Unmarshal(data, sess)
				if ee != nil {
					glog.Debug(ctx, "gwt %s -> %s -> %v : %s", *s, string(data), *sess, ee.Error())
				}
				glog.Debug(ctx, "gwt %s -> %s -> %v ", *s, string(data), *sess)
			}
		}
	}, func(s *string, sess *GSession, gay *gconf.GWebMiddleWareYaml, ctx *gin.Context) {
		sdm, e := json.Marshal(sess)
		glog.Debug(ctx, "shezhi neir %s", string(sdm))
		if e != nil {
			return
		}
		od, e := aes.AesCbcEncryptBase64(sdm, []byte(conf.AuthKey), nil)
		if e != nil {
			return
		}
		*s = od
	})

}
func initRedisClusterMiddleWare(conf *gconf.GWebMiddleWareYaml) gin.HandlerFunc {
	return buildAuthMiddleWare(conf, func(s *string, sess *GSession, gay *gconf.GWebMiddleWareYaml, ctx *gin.Context) {
		if *s == "" {
			*s = uuid.New().String()
		} else {
			data, e := gcache.GetRedisCluster(conf.Ref).HGetAll(ctx, "sess:"+*s).Result()
			if e == nil {
				for k, v := range data {
					sess.Set(k, v)
				}
			}
		}
	}, func(s *string, sess *GSession, gay *gconf.GWebMiddleWareYaml, ctx *gin.Context) {
		vals := []any{}
		for k, v := range *sess {
			vals = append(vals, k, v)
		}
		gcache.GetRedisCluster(conf.Ref).HMSet(ctx, "sess:"+*s, vals...)
		gcache.GetRedisCluster(conf.Ref).Expire(ctx, "sess:"+*s, time.Duration(conf.TimeOut)*time.Second)
	})

}

func initRedisMiddleWare(conf *gconf.GWebMiddleWareYaml) gin.HandlerFunc {
	return buildAuthMiddleWare(conf, func(s *string, sess *GSession, gay *gconf.GWebMiddleWareYaml, ctx *gin.Context) {
		if *s == "" {
			*s = uuid.New().String()
		} else {
			data, e := gcache.GetRedis(conf.Ref).HGetAll(ctx, "sess:"+*s).Result()
			if e == nil {
				for k, v := range data {
					sess.Set(k, v)
				}
			}
		}
	}, func(s *string, sess *GSession, gay *gconf.GWebMiddleWareYaml, ctx *gin.Context) {
		vals := []any{}
		for k, v := range *sess {
			vals = append(vals, k, v)
		}
		gcache.GetRedis(conf.Ref).HMSet(ctx, "sess:"+*s, vals...)
		gcache.GetRedis(conf.Ref).Expire(ctx, "sess:"+*s, time.Duration(conf.TimeOut)*time.Second)
	})

}

func buildAuthMiddleWare(conf *gconf.GWebMiddleWareYaml, befor, after func(*string, *GSession, *gconf.GWebMiddleWareYaml, *gin.Context)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sid := ""
		if conf.IsHeader {
			sid = ctx.GetHeader(conf.SendName)
		}
		if sid == "" && conf.IsCookie {
			sid, _ = ctx.Cookie(conf.SendName)
			sid, _ = url.QueryUnescape(sid)
		}

		glog.Debug(ctx, "read sid %s", sid)
		sess := &GSession{}
		befor(&sid, sess, conf, ctx)

		ctx.Set("session", sess)
		ctx.Next()
		after(&sid, sess, conf, ctx)
		glog.Debug(ctx, "设置 sid %s", sid)
		if conf.IsHeader {
			ctx.Writer.Header().Add(conf.SendName, sid)
			ctx.Header(conf.SendName, sid)
		}
		if conf.IsCookie {
			glog.Debug(ctx, "设置cookie sid %s", sid)
			ctx.SetCookie(conf.SendName, sid, conf.TimeOut, "/", "", false, true)
		}
	}

}
