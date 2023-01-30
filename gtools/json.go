package gtools

import (
	"bytes"
	"fmt"
	"strings"
)

type jsonData struct {
	pkg     map[string]string
	data    map[string]string
	subData map[string]map[string]string
}

func JsonToStruct(src string) string {
	data := jsonData{
		pkg:     map[string]string{},
		data:    map[string]string{},
		subData: map[string]map[string]string{},
	}
	jsonParse(src, &data)
	w := bytes.NewBuffer([]byte{})
	for _, item := range data.pkg {
		w.WriteString(fmt.Sprintf("import \"%s\"\n", item))
	}
	w.WriteString("type AppStruct struct{\n")
	for k, val := range data.data {
		w.WriteString(fmt.Sprintf("%s %s\n", k, val))
	}
	w.WriteString("}\n")
	for k, val := range data.subData {
		w.WriteString(fmt.Sprintf("type %s struct{\n", k))
		for k1, v1 := range val {
			w.WriteString(fmt.Sprintf("%s %s\n", k1, v1))
		}
		w.WriteString("}\n")
	}
	return w.String()
}

func jsonParse(src string, out *jsonData) {
	src = strings.TrimSpace(src)
	sindex := strings.Index(src, "{")
	if sindex == 0 {
		src = strings.TrimSpace(src[:1])
	}
	// eindex := strings.Index(src, "}")

}
