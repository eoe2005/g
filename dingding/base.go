package dingding

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/eoe2005/g/gconf"
	"github.com/eoe2005/g/gnet"
)

const ddapihost = "https://oapi.dingtalk.com/robot/send?access_token="

var (
	_localConf = map[string]*gconf.GDingDingYaml{}
)

// https://open.dingtalk.com/document/robots/custom-robot-access
func Register(confs []*gconf.GDingDingYaml) {
	for _, i := range confs {
		_localConf[i.Name] = i
	}
}

func SendText(ctx context.Context, name, msg string, at ...string) bool {
	sendData := map[string]any{
		"msgtype": "text",
		"text": map[string]string{
			"content": msg,
		},
	}
	if len(at) > 0 {
		sendData["at"] = map[string][]string{
			"atMobiles": at,
		}
	}
	r := send(ctx, name, sendData)
	return r.Err == nil
}
func SendTextAll(ctx context.Context, name, msg string) bool {
	sendData := map[string]any{
		"msgtype": "text",
		"text": map[string]string{
			"content": msg,
		},
		"at": map[string]bool{"isAtAll": true},
	}

	r := send(ctx, name, sendData)
	return r.Err == nil
}
func SendLink(ctx context.Context, name, title, msg, url string, picUrl ...string) bool {
	purl := ""
	if len(picUrl) > 0 {
		purl = picUrl[0]
	}
	sendData := map[string]any{
		"msgtype": "link",
		"link": map[string]string{
			"text":       msg,
			"title":      title,
			"messageUrl": url,
			"picUrl":     purl,
		},
	}

	r := send(ctx, name, sendData)
	return r.Err == nil
}
func SendFeedCard(ctx context.Context, name string, card ...map[string]string) bool {

	sendData := map[string]any{
		"msgtype": "feedCard",
		"feedCard": map[string]any{
			"links": card,
		},
	}

	r := send(ctx, name, sendData)
	return r.Err == nil
}

func SendMarkDown(ctx context.Context, name, title, msg string, at ...string) bool {

	sendData := map[string]any{
		"msgtype": "markdown",
		"markdown": map[string]string{
			"text":  msg,
			"title": title,
		},
	}
	if len(at) > 0 {
		sendData["at"] = map[string][]string{
			"atMobiles": at,
		}
	}
	r := send(ctx, name, sendData)
	return r.Err == nil
}

// https://open.dingtalk.com/document/robots/custom-robot-access
func SendActionCard(ctx context.Context, name string, data any) bool {

	sendData := map[string]any{
		"msgtype":    "actionCard",
		"actionCard": data,
	}

	r := send(ctx, name, sendData)
	return r.Err == nil
}
func SendMarkDownAll(ctx context.Context, name, title, msg string) bool {

	sendData := map[string]any{
		"msgtype": "markdown",
		"markdown": map[string]string{
			"text":  msg,
			"title": title,
		},
		"at": map[string]bool{
			"isAtAll": true,
		},
	}

	r := send(ctx, name, sendData)
	return r.Err == nil
}

func send(ctx context.Context, name string, msg any) *gnet.HttResult {
	url := ""
	if c, ok := _localConf[name]; ok {
		url = ddapihost + c.Token
	}
	if url == "" {
		return &gnet.HttResult{
			Code: -1,
			Err:  errors.New("路径不存在"),
		}
	}
	sd, e := json.Marshal(msg)
	if e != nil {
		return &gnet.HttResult{
			Code: -1,
			Err:  e,
		}
	}
	return gnet.Post(ctx, url, string(sd), 1000, map[string]string{"Content-Type": "application/json;charset=utf-8"})
}
