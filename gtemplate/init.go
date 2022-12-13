package gtemplate

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"io/ioutil"
	"strings"
)

var (
	_template = template.New("")
)

func Fetch(tpl string, data any) string {
	return ""
}

func Register(d embed.FS) {
	ds, e := d.ReadDir("template")
	if e != nil {
		fmt.Println(e)
		return
	}
	for _, f := range ds {
		open(d, f, "template")
	}
}
func open(d embed.FS, ff fs.DirEntry, dir string) {
	dir = dir + "/" + ff.Name()
	if ff.IsDir() {
		ds, e := d.ReadDir(dir)
		if e != nil {
			return
		}
		for _, f := range ds {
			open(d, f, dir)
		}
	} else {
		f, e := d.Open(dir)
		if e != nil {
			return
		}
		defer f.Close()
		data, e := ioutil.ReadAll(f)
		if e != nil {
			return
		}
		dn := strings.Replace(dir, "template/", "", 1)
		dn = strings.ReplaceAll(dn, "/", ".")
		dn = strings.ReplaceAll(dn, ".html", "")
		dn = strings.ReplaceAll(dn, ".htm", "")
		dn = strings.ReplaceAll(dn, ".tpl", "")
		_template.New(dn).Parse(string(data))
	}
}
