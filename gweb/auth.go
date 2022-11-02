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
		data := ""
		if conf.IsHeader {
			data = ctx.GetHeader(conf.SendName)
		}
		if data == "" && conf.IsCookie {
			data, _ = ctx.Cookie(conf.SendName)
		}
		sess := GSession{}
		if data != "" {
			data, err := aes.AesCbcDecryptByBase64(data, []byte(conf.AuthKey), nil)
			if err == nil {
				json.Unmarshal(data, sess)
			}

		}
		ctx.Set("session", sess)
		ctx.Next()

		sdm, e := json.Marshal(sess)
		if e != nil {
			return
		}
		od, e := aes.AesCbcEncryptBase64(sdm, []byte(conf.AuthKey), nil)
		if e != nil {
			return
		}
		if conf.IsHeader {
			ctx.Writer.Header().Add(conf.SendName, od)
		}
		if conf.IsCookie {
			ctx.SetCookie(conf.SendName, od, conf.TimeOut, "/", "", false, false)
		}
	}

}
func getRedisClusterMiddleWare(conf *gconf.GWebAuthYaml) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sid := ""
		if conf.IsHeader {
			sid = ctx.GetHeader(conf.SendName)
		}
		if sid == "" && conf.IsCookie {
			sid, _ = ctx.Cookie(conf.SendName)
		}
		sess := GSession{}
		if sid != "" {
			data, e := gcache.GetRedisCluster(conf.Ref).HGetAll(ctx, "sess:"+sid).Result()
			if e == nil {
				for k, v := range data {
					sess.Set(k, v)
				}
			}
		} else {
			sid = uuid.New().String()
		}
		ctx.Set("session", sess)
		ctx.Next()
		vals := []any{}
		for k, v := range sess {
			vals = append(vals, k, v)
		}
		gcache.GetRedisCluster(conf.Ref).HMSet(ctx, "sess:"+sid, vals...)
		gcache.GetRedisCluster(conf.Ref).Expire(ctx, "sess:"+sid, time.Duration(conf.TimeOut)*time.Second)
		if conf.IsHeader {
			ctx.Writer.Header().Add(conf.SendName, sid)
		}
		if conf.IsCookie {
			ctx.SetCookie(conf.SendName, sid, conf.TimeOut, "/", "", false, false)
		}
	}

}

func getRedisMiddleWare(conf *gconf.GWebAuthYaml) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sid := ""
		if conf.IsHeader {
			sid = ctx.GetHeader(conf.SendName)
		}
		if sid == "" && conf.IsCookie {
			sid, _ = ctx.Cookie(conf.SendName)
		}
		sess := GSession{}
		if sid != "" {
			data, e := gcache.GetRedis(conf.Ref).HGetAll(ctx, "sess:"+sid).Result()
			if e == nil {
				for k, v := range data {
					sess.Set(k, v)
				}
			}
		} else {
			sid = uuid.New().String()
		}
		ctx.Set("session", sess)
		ctx.Next()
		vals := []any{}
		for k, v := range sess {
			vals = append(vals, k, v)
		}
		gcache.GetRedis(conf.Ref).HMSet(ctx, "sess:"+sid, vals...)
		gcache.GetRedis(conf.Ref).Expire(ctx, "sess:"+sid, time.Duration(conf.TimeOut)*time.Second)
		if conf.IsHeader {
			ctx.Writer.Header().Add(conf.SendName, sid)
		}
		if conf.IsCookie {
			ctx.SetCookie(conf.SendName, sid, conf.TimeOut, "/", "", false, false)
		}
	}
}
