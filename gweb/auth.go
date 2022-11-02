package gweb

import (
	"encoding/json"
	"time"

	"github.com/eoe2005/g/gcache"
	"github.com/eoe2005/g/gconf"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wumansgy/goEncrypt/aes"
)

func getAuthMiddleWare(conf *gconf.GWebAuthYaml) gin.HandlerFunc {
	switch conf.Driver {
	case "jwt":
		return getJwtMiddleWare(conf)
	case "redis":
		return getRedisMiddleWare(conf)
	case "redis_cluster":
		return getRedisClusterMiddleWare(conf)
	}
	return nil
}
func getJwtMiddleWare(conf *gconf.GWebAuthYaml) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		buildAuthMiddleWare(conf, ctx, func(s *string, sess GSession, gay *gconf.GWebAuthYaml, ctx *gin.Context) {
			if *s != "" {
				data, err := aes.AesCbcDecryptByBase64(*s, []byte(conf.AuthKey), nil)
				if err == nil {
					json.Unmarshal(data, sess)
				}
			}
		}, func(s *string, sess GSession, gay *gconf.GWebAuthYaml, ctx *gin.Context) {
			sdm, e := json.Marshal(sess)
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

}
func getRedisClusterMiddleWare(conf *gconf.GWebAuthYaml) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		buildAuthMiddleWare(conf, ctx, func(s *string, sess GSession, gay *gconf.GWebAuthYaml, ctx *gin.Context) {
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
		}, func(s *string, sess GSession, gay *gconf.GWebAuthYaml, ctx *gin.Context) {
			vals := []any{}
			for k, v := range sess {
				vals = append(vals, k, v)
			}
			gcache.GetRedisCluster(conf.Ref).HMSet(ctx, "sess:"+*s, vals...)
			gcache.GetRedisCluster(conf.Ref).Expire(ctx, "sess:"+*s, time.Duration(conf.TimeOut)*time.Second)
		})

	}

}

func getRedisMiddleWare(conf *gconf.GWebAuthYaml) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		buildAuthMiddleWare(conf, ctx, func(s *string, sess GSession, gay *gconf.GWebAuthYaml, ctx *gin.Context) {
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
		}, func(s *string, sess GSession, gay *gconf.GWebAuthYaml, ctx *gin.Context) {
			vals := []any{}
			for k, v := range sess {
				vals = append(vals, k, v)
			}
			gcache.GetRedis(conf.Ref).HMSet(ctx, "sess:"+*s, vals...)
			gcache.GetRedis(conf.Ref).Expire(ctx, "sess:"+*s, time.Duration(conf.TimeOut)*time.Second)
		})

	}
}

func buildAuthMiddleWare(conf *gconf.GWebAuthYaml, ctx *gin.Context, befor, after func(*string, GSession, *gconf.GWebAuthYaml, *gin.Context)) {
	sid := ""
	if conf.IsHeader {
		sid = ctx.GetHeader(conf.SendName)
	}
	if sid == "" && conf.IsCookie {
		sid, _ = ctx.Cookie(conf.SendName)
	}
	sess := GSession{}
	befor(&sid, sess, conf, ctx)

	ctx.Set("session", sess)
	ctx.Next()
	after(&sid, sess, conf, ctx)

	if conf.IsHeader {
		ctx.Writer.Header().Add(conf.SendName, sid)
	}
	if conf.IsCookie {
		ctx.SetCookie(conf.SendName, sid, conf.TimeOut, "/", "", false, false)
	}
}
