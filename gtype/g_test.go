package gtype_test

import (
	"testing"

	"github.com/eoe2005/g/gtype"
)

func TestFetchFeildsByName(t *testing.T) {
	dataS := []struct {
		Name  string `json:"t1"`
		Name2 string `json:"t2"`
	}{
		{Name: "test1", Name2: "weqrwerq"}, {Name: "nihao"},
	}

	dataM := []struct {
		Name  string `json:"t1"`
		Name2 string `json:"t2"`
	}{
		{Name: "test1", Name2: "weqrwerq"}, {Name: "nihao"},
	}
	r1 := gtype.FetchFeildsByName(dataS, "t1", "t2")
	r2 := gtype.FetchFeildsByName(dataM, "t1", "t2")
	r11 := gtype.FetchFeildsByName(&dataS, "t1", "t2")
	r22 := gtype.FetchFeildsByName(&dataM, "t1", "t2")
	t.Logf("ds-> %v\nr1 -> %v\nr2 -> %v\nr11 -> %v\n22 -> %v\n", dataS, r1, r2, r11, r22)
	// for _, item := range r1 {
	// 	item["t1"].SetString("你好")
	// }
	// t.Logf("ds-> %v\nr1 -> %v\nr2 -> %v\nr11 -> %v\n22 -> %v\n", dataS, r1, r2, r11, r22)
	// for _, item := range r11 {
	// 	item["t1"].SetString("你好")
	// }
	// t.Logf("ds-> %v\nr1 -> %v\nr2 -> %v\nr11 -> %v\n22 -> %v\n", dataS, r1, r2, r11, r22)
}
