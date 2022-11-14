package gweb

import (
	"encoding/json"
	"reflect"

	"github.com/gin-gonic/gin"
)

func BindJson(c *gin.Context, obj any) error {
	ret := json.NewDecoder(c.Request.Body).Decode(obj)
	if ret != nil {
		return ret
	}
	v := reflect.ValueOf(obj).Elem().FieldByName("UserID")
	if v.CanUint() && v.Uint() == 0 {
		v.SetUint(uint64(GetSession(c).GetInt("userid", 0)))
	} else if v.CanInt() && v.Int() == 0 {
		v.SetInt(GetSession(c).GetInt("userid", 0))
	}
	v2 := reflect.ValueOf(obj).Elem().FieldByName("AdminUID")
	if v2.CanUint() && v2.Uint() == 0 {
		v2.SetUint(uint64(GetSession(c).GetInt("adminuid", 0)))
	} else if v2.CanInt() && v2.Int() == 0 {
		v2.SetInt(GetSession(c).GetInt("adminuid", 0))
	}
	return nil
}
