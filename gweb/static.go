package gweb

import (
	"embed"
	"strings"

	"github.com/gin-gonic/gin"
)

func StaticFs(c *gin.Context, path, fspath string, f embed.FS) {
	filename := strings.TrimLeft(c.Request.URL.Path, path+"/")
	if filename == "" {
		filename = "index.html"
	}
	// fmt.Println("1", fspath+filename)
	data, e := f.ReadFile(fspath + filename)
	if e != nil {
		c.Abort()
	}
	// fmt.Println("sendData", fspath+filename)
	if strings.HasSuffix(filename, ".js") {
		c.Writer.Header().Add("Content-Type", "application/javascript")
	} else if strings.HasSuffix(filename, ".css") {
		c.Writer.Header().Add("Content-Type", "text/css")
	} else if strings.HasSuffix(filename, ".gif") {
		c.Writer.Header().Add("Content-Type", "image/gif")
	} else if strings.HasSuffix(filename, ".png") {
		c.Writer.Header().Add("Content-Type", "image/png")
	} else if strings.HasSuffix(filename, ".jpg") {
		c.Writer.Header().Add("Content-Type", "image/jpg")
	} else if strings.HasSuffix(filename, ".jpeg") {
		c.Writer.Header().Add("Content-Type", "image/jpg")
	}
	c.Writer.Write(data)
	c.Abort()

}
