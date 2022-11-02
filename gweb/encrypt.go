package gweb

import (
	"bytes"

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
func getEncryptMiddleWare(conf *gconf.GWebEncryptYaml) gin.HandlerFunc {
	switch conf.Driver {
	case "aes":
		return getAesMiddleWare(conf)
	case "rsa":
		return getSslMiddleWare(conf)
	}
	return nil
}
func getAesMiddleWare(conf *gconf.GWebEncryptYaml) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		oldWriter := ctx.Writer
		blw := &gineEncryptWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = blw
		ctx.Next()
		responseByte := blw.body.Bytes()

		ctx.Writer = oldWriter
		aesSecretKey := conf.Key
		base64Text, _ := aes.AesCbcEncryptBase64([]byte(responseByte), []byte(aesSecretKey), nil)
		ctx.Writer.WriteString(base64Text)
	}
}

func getSslMiddleWare(conf *gconf.GWebEncryptYaml) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		oldWriter := ctx.Writer
		blw := &gineEncryptWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = blw
		ctx.Next()
		responseByte := blw.body.Bytes()

		ctx.Writer = oldWriter
		rsa.GenerateRsaKeyBase64(1024)
		base64Text, _ := rsa.RsaEncryptToBase64(responseByte, conf.PublicKey)

		ctx.Writer.WriteString(base64Text)
	}
}
