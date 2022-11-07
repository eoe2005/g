package gweb

import (
	"bytes"
	"io/ioutil"

	"github.com/eoe2005/g/gconf"
	"github.com/gin-gonic/gin"
	"github.com/wumansgy/goEncrypt/aes"
	"github.com/wumansgy/goEncrypt/rsa"
)

type gineEncryptWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *gineEncryptWriter) Write(b []byte) (int, error) {
	return w.body.Write(b)
}
func (w *gineEncryptWriter) WriteString(s string) (int, error) {
	return w.body.WriteString(s)
}

func initAesMiddleWare(conf *gconf.GWebMiddleWareYaml) gin.HandlerFunc {
	return resetInput(conf, func(b []byte, gey *gconf.GWebMiddleWareYaml) []byte {
		r, _ := aes.AesCbcDecryptByBase64(string(b), []byte(gey.Key), nil)
		return r
	}, func(b []byte, gey *gconf.GWebMiddleWareYaml) string {
		r, _ := aes.AesCbcEncryptBase64([]byte(b), []byte(gey.Key), nil)
		return r
	})

}

func initRsaMiddleWare(conf *gconf.GWebMiddleWareYaml) gin.HandlerFunc {
	return resetInput(conf, func(b []byte, gey *gconf.GWebMiddleWareYaml) []byte {
		r, _ := rsa.RsaDecryptByBase64(string(b), gey.PrivateKey)
		return r
	}, func(b []byte, gey *gconf.GWebMiddleWareYaml) string {
		r, _ := rsa.RsaEncryptToBase64(b, gey.PublicKey)
		return r
	})

}

func resetInput(conf *gconf.GWebMiddleWareYaml, before func([]byte, *gconf.GWebMiddleWareYaml) []byte, after func([]byte, *gconf.GWebMiddleWareYaml) string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, e := ioutil.ReadAll(ctx.Request.Body)
		if e != nil {
			idata := before(data, conf)
			ctx.Request.Write(bytes.NewBuffer(idata))
		}
		oldWriter := ctx.Writer
		blw := &gineEncryptWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = blw
		ctx.Next()
		responseByte := blw.body.Bytes()

		ctx.Writer = oldWriter
		base64Text := after(responseByte, conf)

		ctx.Writer.WriteString(base64Text)
	}

}
