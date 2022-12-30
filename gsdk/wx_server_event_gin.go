package gsdk

import (
	"strings"

	"github.com/eoe2005/g/glog"
	"github.com/gin-gonic/gin"
)

type WxServerCallBackHandle func(*gin.Context, WxServerEvent)
type WxServerCallBackFuncStruct struct {
	MsgTextHandle              WxServerCallBackHandle
	MsgImageHandle             WxServerCallBackHandle
	MsgVoiceHandle             WxServerCallBackHandle
	MsgVideoHandle             WxServerCallBackHandle
	MsgLocationHandle          WxServerCallBackHandle
	MsgShortVideoHandle        WxServerCallBackHandle
	MsgLinkHandle              WxServerCallBackHandle
	EventSubscribeHandle       WxServerCallBackHandle
	EventSubscribeScanHandle   WxServerCallBackHandle
	EvnetUnSubscribeHandle     WxServerCallBackHandle
	EventScanHandle            WxServerCallBackHandle
	EventLocationHandle        WxServerCallBackHandle
	EventClickHandle           WxServerCallBackHandle
	EventViewHandle            WxServerCallBackHandle
	EventScanCodePushHandle    WxServerCallBackHandle
	EventScanCodeWaitmsgHandle WxServerCallBackHandle
	EventPicSysPhotoHandle     WxServerCallBackHandle
	EventPicPhotoOrAlbumHandle WxServerCallBackHandle
	EventPicWeixinHandle       WxServerCallBackHandle
	EventLocationSelectHandle  WxServerCallBackHandle
	EventViewMiniProgramHandle WxServerCallBackHandle
}

func WxServerCallBack(c *gin.Context, call *WxServerCallBackFuncStruct) {
	if call == nil {
		return
	}
	req := WxServerEvent{}
	e := c.BindXML(&req)
	if e != nil {
		glog.Error(c, "读取是时间失败 %s", e.Error())
		return
	}
	switch req.MsgType {
	case "text": //文本消息
		if call.MsgTextHandle != nil {
			call.MsgTextHandle(c, req)
		}
	case "image": //图片消息
		if call.MsgImageHandle != nil {
			call.MsgImageHandle(c, req)
		}
	case "voice": //语音消息
		if call.MsgVoiceHandle != nil {
			call.MsgVoiceHandle(c, req)
		}
	case "video": //视频消息
		if call.MsgVideoHandle != nil {
			call.MsgVideoHandle(c, req)
		}
	case "shortvideo": //小视频消息
		if call.MsgShortVideoHandle != nil {
			call.MsgShortVideoHandle(c, req)
		}
	case "location": //地理位置消息
		if call.MsgLocationHandle != nil {
			call.MsgLocationHandle(c, req)
		}
	case "link": //链接消息
		if call.MsgLinkHandle != nil {
			call.MsgLinkHandle(c, req)
		}
	case "event":
		switch req.Event { //subscribe(订阅)、unsubscribe(取消订阅)
		case "subscribe": //订阅
			if strings.HasPrefix(req.EventKey, "qrscene_") { //扫描带参数二维码事件
				if call.EventSubscribeScanHandle != nil {
					call.EventSubscribeScanHandle(c, req)
				}
			} else {
				if call.EventSubscribeHandle != nil {
					call.EventSubscribeHandle(c, req)
				}
			}
		case "unsubscribe": //取消订阅
			if call.EvnetUnSubscribeHandle != nil {
				call.EvnetUnSubscribeHandle(c, req)
			}
		case "SCAN": //用户已关注时的事件推送
			if call.EventScanHandle != nil {
				call.EventScanHandle(c, req)
			}
		case "LOCATION": //上报地理位置事件
			if call.EventLocationHandle != nil {
				call.EventLocationHandle(c, req)
			}
		case "CLICK": //自定义菜单事件
			if req.EventKey == "EVENTKEY" { //点击菜单拉取消息时的事件推送

			}
			if call.EventClickHandle != nil {
				call.EventClickHandle(c, req)
			}
		case "VIEW": //点击菜单跳转链接时的事件推送
			if call.EventViewHandle != nil {
				call.EventViewHandle(c, req)
			}
		case "scancode_push": //scancode_push：扫码推事件的事件推送
			if call.EventScanCodePushHandle != nil {
				call.EventScanCodePushHandle(c, req)
			}
		case "scancode_waitmsg": //scancode_waitmsg：扫码推事件且弹出“消息接收中”提示框的事件推送
			if call.EventScanCodeWaitmsgHandle != nil {
				call.EventScanCodeWaitmsgHandle(c, req)
			}
		case "pic_sysphoto": //pic_sysphoto：弹出系统拍照发图的事件推送
			if call.EventPicSysPhotoHandle != nil {
				call.EventPicSysPhotoHandle(c, req)
			}
		case "pic_photo_or_album": //pic_photo_or_album：弹出拍照或者相册发图的事件推送
			if call.EventPicPhotoOrAlbumHandle != nil {
				call.EventPicPhotoOrAlbumHandle(c, req)
			}
		case "pic_weixin": //pic_weixin：弹出微信相册发图器的事件推送
			if call.EventPicWeixinHandle != nil {
				call.EventPicWeixinHandle(c, req)
			}
		case "location_select": //location_select：弹出地理位置选择器的事件推送
			if call.EventLocationSelectHandle != nil {
				call.EventLocationSelectHandle(c, req)
			}
		case "view_miniprogram": //点击菜单跳转小程序的事件推送
			if call.EventViewMiniProgramHandle != nil {
				call.EventViewMiniProgramHandle(c, req)
			}
		}
	}
}
