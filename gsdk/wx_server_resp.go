package gsdk

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

//回复文本消息
func WxServerSendText(c *gin.Context, in WxServerEvent, msg string) {
	c.Writer.Write([]byte(fmt.Sprintf(`<xml>
	<ToUserName><![CDATA[%s]]></ToUserName>
	<FromUserName><![CDATA[%s]]></FromUserName>
	<CreateTime>%d</CreateTime>
	<MsgType><![CDATA[text]]></MsgType>
	<Content><![CDATA[%s]]></Content>
  </xml>`, in.FromUserName, in.ToUserName, time.Now().Unix(), msg)))

}

//回复图片消息
func WxServerSendImage(c *gin.Context, in WxServerEvent, MediaId string) {
	c.Writer.Write([]byte(fmt.Sprintf(`<xml>
	<ToUserName><![CDATA[%s]]></ToUserName>
	<FromUserName><![CDATA[%s]]></FromUserName>
	<CreateTime>%d</CreateTime>
	<MsgType><![CDATA[image]]></MsgType>
	<Image>
	  <MediaId><![CDATA[%s]]></MediaId>
	</Image>
  </xml>`, in.FromUserName, in.ToUserName, time.Now().Unix(), MediaId)))

}

//回复语音消息
func WxServerSendVoice(c *gin.Context, in WxServerEvent, MediaId string) {
	c.Writer.Write([]byte(fmt.Sprintf(`<xml>
	<ToUserName><![CDATA[%s]]></ToUserName>
	<FromUserName><![CDATA[%s]]></FromUserName>
	<CreateTime>%d</CreateTime>
	<MsgType><![CDATA[voice]]></MsgType>
	<Voice>
	  <MediaId><![CDATA[%s]]></MediaId>
	</Voice>
  </xml>`, in.FromUserName, in.ToUserName, time.Now().Unix(), MediaId)))

}

//回复视频消息
func WxServerSendVideo(c *gin.Context, in WxServerEvent, MediaId, title, desc string) {
	c.Writer.Write([]byte(fmt.Sprintf(`<xml>
	<ToUserName><![CDATA[%s]]></ToUserName>
	<FromUserName><![CDATA[%s]]></FromUserName>
	<CreateTime>%d</CreateTime>
	<MsgType><![CDATA[video]]></MsgType>
	<Video>
	  <MediaId><![CDATA[%s]]></MediaId>
	  <Title><![CDATA[%s]]></Title>
	  <Description><![CDATA[%s]]></Description>
	</Video>
  </xml>`, in.FromUserName, in.ToUserName, time.Now().Unix(), MediaId, title, desc)))

}

//回复音乐消息
func WxServerSendMusic(c *gin.Context, in WxServerEvent, Title, Description, MusicUrl, HQMusicUrl, ThumbMediaId string) {
	c.Writer.Write([]byte(fmt.Sprintf(`<xml>
	<ToUserName><![CDATA[%s]]></ToUserName>
	<FromUserName><![CDATA[%s]]></FromUserName>
	<CreateTime>%d</CreateTime>
	<MsgType><![CDATA[music]]></MsgType>
	<Music>
	  <ThumbMediaId><![CDATA[%s]]></ThumbMediaId>
	  <Title><![CDATA[%s]]></Title>
	  <Description><![CDATA[%s]]></Description>
	  <MusicUrl><![CDATA[%s]]></MusicUrl>
	<HQMusicUrl><![CDATA[%s]]></HQMusicUrl>
	  
	</Music>
  </xml>`, in.FromUserName, in.ToUserName, time.Now().Unix(), ThumbMediaId, Title, Description, MusicUrl, HQMusicUrl)))

}

type WxServerSendArticlesStructList []WxServerSendArticlesStruct
type WxServerSendArticlesStruct struct {
	Title       string
	Description string
	PicUrl      string
	Url         string
}

//回复图文消息
func WxServerSendArticles(c *gin.Context, in WxServerEvent, data WxServerSendArticlesStructList) {

	str := ""
	for _, item := range data {
		str += fmt.Sprintf(`<item>
		<Title><![CDATA[%s]]></Title>
		<Description><![CDATA[%s]]></Description>
		<PicUrl><![CDATA[%s]]></PicUrl>
		<Url><![CDATA[%s]]></Url>
	  </item>`, item.Title, item.Description, item.PicUrl, item.Url)
	}
	c.Writer.Write([]byte(fmt.Sprintf(`<xml>
	<ToUserName><![CDATA[%s]]></ToUserName>
	<FromUserName><![CDATA[%s]]></FromUserName>
	<CreateTime>%d</CreateTime>
	<MsgType><![CDATA[news]]></MsgType>
	<ArticleCount>%d</ArticleCount>
	<Articles>
	  %s
	</Articles>
  </xml>`, in.FromUserName, in.ToUserName, time.Now().Unix(), len(data), str)))

}
