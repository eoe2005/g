package output

import (
	"embed"
	"html/template"

	"github.com/gin-gonic/gin"
)

var (
	_tpl = template.New("g")
)

func Html(c *gin.Context, name string, data any) {
	_tpl.ExecuteTemplate(c.Writer, name, data)
	c.Abort()
}

func SetFs(prefix string, fs embed.FS) {
	ds, e := fs.ReadDir(".")
	if e != nil {
		return
	}
	for _, f := range ds {
		if f.IsDir() {
		} else {
			ts, e := fs.ReadFile(f.Name())
			if e != nil {
				continue
			}
			_tpl.New(prefix + "." + f.Name()).Parse(string(ts))
		}
	}
}
