package gweb

import (
	"embed"
	"strings"

	"github.com/gin-gonic/gin"
)

//注册vue
func VueHistory(e *gin.Engine, path, fspath string, f embed.FS) {
	e.GET(path+"/*filepath", func(ctx *gin.Context) {
		_vueHistory(path, fspath, f, ctx)
	})
	e.HEAD(path+"/*filepath", func(ctx *gin.Context) {
		_vueHistory(path, fspath, f, ctx)
	})

}

//注册vue
func VueHistoryGroup(group *gin.RouterGroup, path, fspath string, f embed.FS) {
	group.GET(path+"/*filepath", func(ctx *gin.Context) {
		_vueHistory(group.BasePath()+path, fspath, f, ctx)
	})
	group.HEAD(path+"/*filepath", func(ctx *gin.Context) {
		_vueHistory(group.BasePath()+path, fspath, f, ctx)
	})

}

func _vueHistory(path, fspath string, f embed.FS, c *gin.Context) {
	filename := strings.TrimLeft(c.Request.URL.Path, path+"/")
	if filename == "" {
		filename = "index.html"
	}
	for {
		// fmt.Println("1", fspath+filename)
		data, e := f.ReadFile(fspath + filename)
		if e != nil {
			if filename == "index.html" {
				// fmt.Println("文件不存在", fspath+filename)
				// fmt.Fprintln("文件不存在", fspath+filename)
				return
			} else {
				filename = "index.html"
				// fmt.Println("文件不存在2", filename)
			}
		}
		// fmt.Println("sendData", fspath+filename)
		if strings.HasSuffix(filename, ".js") {
			c.Writer.Header().Add("Content-Type", "application/javascript")
		} else if strings.HasSuffix(filename, ".css") {
			c.Writer.Header().Add("Content-Type", "text/css")
		}
		c.Writer.Write(data)
		c.Abort()
		return
	}

}
