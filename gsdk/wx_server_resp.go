package gsdk

import (
	"github.com/gin-gonic/gin"
)

//回复文本消息
func WxServerSendText(c *gin.Context, in WxServerEvent, msg string) {
	// 	<xml>
	//   <ToUserName><![CDATA[toUser]]></ToUserName>
	//   <FromUserName><![CDATA[fromUser]]></FromUserName>
	//   <CreateTime>12345678</CreateTime>
	//   <MsgType><![CDATA[text]]></MsgType>
	//   <Content><![CDATA[你好]]></Content>
	// </xml>
}

//回复图片消息
func WxServerSendImage(c *gin.Context, in WxServerEvent, MediaId string) {
	// 	<xml>
	//   <ToUserName><![CDATA[toUser]]></ToUserName>
	//   <FromUserName><![CDATA[fromUser]]></FromUserName>
	//   <CreateTime>12345678</CreateTime>
	//   <MsgType><![CDATA[image]]></MsgType>
	//   <Image>
	//     <MediaId><![CDATA[media_id]]></MediaId>
	//   </Image>
	// </xml>
}

//回复语音消息
func WxServerSendVoice(c *gin.Context, in WxServerEvent, MediaId string) {
	// 	<xml>
	//   <ToUserName><![CDATA[toUser]]></ToUserName>
	//   <FromUserName><![CDATA[fromUser]]></FromUserName>
	//   <CreateTime>12345678</CreateTime>
	//   <MsgType><![CDATA[voice]]></MsgType>
	//   <Voice>
	//     <MediaId><![CDATA[media_id]]></MediaId>
	//   </Voice>
	// </xml>
}

//回复视频消息
func WxServerSendVideo(c *gin.Context, in WxServerEvent, MediaId, title, desc string) {
	// 	<xml>
	// 	<ToUserName><![CDATA[toUser]]></ToUserName>
	// 	<FromUserName><![CDATA[fromUser]]></FromUserName>
	// 	<CreateTime>12345678</CreateTime>
	// 	<MsgType><![CDATA[video]]></MsgType>
	// 	<Video>
	// 	  <MediaId><![CDATA[media_id]]></MediaId>
	// 	  <Title><![CDATA[title]]></Title>
	// 	  <Description><![CDATA[description]]></Description>
	// 	</Video>
	//   </xml>
}

//回复音乐消息
func WxServerSendMusic(c *gin.Context, in WxServerEvent, Title, Description, MusicUrl, HQMusicUrl, ThumbMediaId string) {
	// 	<xml>
	// 	<ToUserName><![CDATA[toUser]]></ToUserName>
	// 	<FromUserName><![CDATA[fromUser]]></FromUserName>
	// 	<CreateTime>12345678</CreateTime>
	// 	<MsgType><![CDATA[music]]></MsgType>
	// 	<Music>
	// 	  <Title><![CDATA[TITLE]]></Title>
	// 	  <Description><![CDATA[DESCRIPTION]]></Description>
	// 	  <MusicUrl><![CDATA[MUSIC_Url]]></MusicUrl>
	// 	  <HQMusicUrl><![CDATA[HQ_MUSIC_Url]]></HQMusicUrl>
	// 	  <ThumbMediaId><![CDATA[media_id]]></ThumbMediaId>
	// 	</Music>
	//   </xml>
}

//回复图文消息
func WxServerSendArticles(c *gin.Context, in WxServerEvent, data []map[string]string) {
	// <xml>
	// <ToUserName><![CDATA[toUser]]></ToUserName>
	// <FromUserName><![CDATA[fromUser]]></FromUserName>
	// <CreateTime>12345678</CreateTime>
	// <MsgType><![CDATA[news]]></MsgType>
	// <ArticleCount>1</ArticleCount>
	// <Articles>
	//   <item>
	// 	<Title><![CDATA[title1]]></Title>
	// 	<Description><![CDATA[description1]]></Description>
	// 	<PicUrl><![CDATA[picurl]]></PicUrl>
	// 	<Url><![CDATA[url]]></Url>
	//   </item>
	// </Articles
}
