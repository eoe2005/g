package gsdk

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/eoe2005/g/glog"
	"github.com/eoe2005/g/gnet"
	"github.com/eoe2005/g/middle"
)

type WxServerTokenResp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
}
type WxServerBaseResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func GetWxServerToken(ctx context.Context, key string) string {
	ret := middle.Handle(middle.WX_SERVER_TOKEN, ctx, key, "")
	token, ok := ret.(string)
	if ok {
		return token
	}
	return ""
}

// 获取服务器的Token
// https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_access_token.html
func WxServerToken(c context.Context, name string) WxServerTokenResp {
	conf := getWxConf(name)
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", conf.AppID, conf.AppSecret)
	ret := gnet.Get(c, url, 3, map[string]string{})
	if ret.Err != nil {
		return WxServerTokenResp{}
	} else {
		r := WxServerTokenResp{}
		if json.Unmarshal(ret.Body, &r) != nil {
			return WxServerTokenResp{}
		}
		return r
	}
}

// 定时的或更新wx服务端的信息
func WxServerTokenLoop(c context.Context, name string, call func(context.Context, string, string)) {
	ticker := time.NewTicker(7100 * time.Second)
	token := WxServerToken(c, name)
	if token.AccessToken != "" {
		call(c, name, token.AccessToken)
	}
	go func() {
		for {
			//从定时器中获取数据
			<-ticker.C
			token := WxServerToken(c, name)
			if token.AccessToken != "" {
				call(c, name, token.AccessToken)
			}

		}
	}()

}

// 更新菜单
func WxServerPushMenu(c context.Context, token string, data []byte) bool {
	ret := gnet.Post(c, fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/menu/create?access_token=%s", token), string(data), 3, map[string]string{})
	if ret.Err != nil {
		glog.Error(c, "更新菜单失败 -> %s", ret.Err.Error())
		return false
	} else {
		r := WxServerBaseResp{}
		if json.Unmarshal(ret.Body, &r) != nil {
			glog.Error(c, "更新菜单失败 -> %s", ret.Err.Error())
			return false
		}
		return r.Errcode == 0
	}
}
