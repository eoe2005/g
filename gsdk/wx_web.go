package gsdk

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/eoe2005/g/gconf"
	"github.com/eoe2005/g/gnet"
)

var (
	wxconfMap = map[string]*gconf.WxYaml{}
)

type WxAccessTokenResp struct {
	ConfName       string
	AccessToken    string `json:"access_token"`
	Errcode        int    `json:"errcode"`
	Errmsg         string `json:"errmsg"`
	ExpiresIn      int    `json:"expires_in"`
	IsSnapshotuser int    `json:"is_snapshotuser"`
	Openid         string `json:"openid"`
	RefreshToken   string `json:"refresh_token"`
	Scope          string `json:"scope"`
	Unionid        string `json:"unionid"`
}
type WxWebRefreshTokenResp struct {
	ConfName     string
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	Openid       string `json:"openid"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	Errcode      int    `json:"errcode"`
	Errmsg       string `json:"errmsg"`
}
type WxWebUserInfo struct {
	Openid     string   `json:"openid"`
	Nickname   string   `json:"nickname"`
	Sex        int      `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	Headimgurl string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	Unionid    string   `json:"unionid"`
	Errcode    int      `json:"errcode"`
	Errmsg     string   `json:"errmsg"`
}

func RegisterWxConf(lst []*gconf.WxYaml) {
	for _, item := range lst {
		wxconfMap[item.Name] = item
	}
}
func getWxConf(name string) *gconf.WxYaml {
	ret, ok := wxconfMap[name]
	if ok {
		return ret
	}
	return &gconf.WxYaml{}
}

//web 登录跳转
func WxWebRedirect(name string) string {
	conf := getWxConf(name)
	if conf.Scope != "snsapi_userinfo" {
		conf.Scope = "snsapi_base"
	}
	return fmt.Sprintf("https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=STATE#wechat_redirect", conf.AppID, url.QueryEscape(conf.RetureUrl), conf.Scope)

}

//微信网页板获取登录的accession token信息
func WxWebGetAccessToken(c context.Context, name, code string) WxAccessTokenResp {
	conf := getWxConf(name)
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", conf.AppID, conf.AppSecret, code)
	ret := gnet.Get(c, url, 3, map[string]string{})
	if ret.Err != nil {
		return WxAccessTokenResp{
			ConfName: name,
		}
	} else {
		r := WxAccessTokenResp{ConfName: name}
		if json.Unmarshal(ret.Body, &r) != nil {
			return WxAccessTokenResp{ConfName: name}
		}
		return r
	}
}

//刷线token
func (wr WxAccessTokenResp) Refresh(c context.Context) WxWebRefreshTokenResp {
	conf := getWxConf(wr.ConfName)
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%s&grant_type=refresh_token&refresh_token=%s", conf.AppID, wr.RefreshToken)
	ret := gnet.Get(c, url, 3, map[string]string{})
	if ret.Err != nil {
		return WxWebRefreshTokenResp{
			ConfName: wr.ConfName,
		}
	} else {
		r := WxWebRefreshTokenResp{ConfName: wr.ConfName}
		if json.Unmarshal(ret.Body, &r) != nil {
			return WxWebRefreshTokenResp{ConfName: wr.ConfName}
		}
		return r
	}
}

//获取微信用户账号信息
func WxWebGetUserInfo(c context.Context, access_token, openid string) WxWebUserInfo {
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN", access_token, openid)
	ret := gnet.Get(c, url, 3, map[string]string{})
	if ret.Err != nil {
		return WxWebUserInfo{}
	} else {
		r := WxWebUserInfo{}
		if json.Unmarshal(ret.Body, &r) != nil {
			return WxWebUserInfo{}
		}
		return r
	}

}
