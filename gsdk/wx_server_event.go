package gsdk

import "encoding/xml"

//消息类型
type WxServerEvent struct {
	Name             xml.Name                      `xml:"xml"`
	ToUserName       string                        `xml:"ToUserName"`
	FromUserName     string                        `xml:"FromUserName"`
	CreateTime       int64                         `xml:"CreateTime"`
	MsgType          string                        `xml:"MsgType"`
	Event            string                        `xml:"Event"`
	EventKey         string                        `xml:"EventKey"`
	MenuId           string                        `xml:"MenuId"`
	Ticket           string                        `xml:"Ticket"`
	ScanCodeInfo     *WxServerEventScanCodeInfo    `xml:"ScanCodeInfo"`
	SendPicsInfo     *WxServerEventSendPicsInfo    `xml:"SendPicsInfo"`
	SendLocationInfo WxServerEventSendLocationInfo `xml:"SendLocationInfo"`
	Latitude         float64                       `xml:"Latitude"`
	Longitude        float64                       `xml:"Longitude"`
	Precision        float64                       `xml:"Precision"`
	Content          string                        `xml:"Content"`
	MsgId            int64                         `xml:"MsgId"`
	MsgDataId        string                        `xml:"MsgDataId"`
	Idx              string                        `xml:"Idx"`
	PicUrl           string                        `xml:"PicUrl"`
	MediaId          string                        `xml:"MediaId"`
	Format           string                        `xml:"Format"`
	Recognition      string                        `xml:"Recognition"`
	ThumbMediaId     string                        `xml:"ThumbMediaId"`
	Location_X       float64                       `xml:"Location_X"`
	Location_Y       float64                       `xml:"Location_Y"`
	Scale            int                           `xml:"Scale"`
	Label            string                        `xml:"Label"`

	Title       string `xml:"Title"`
	Description string `xml:"Description"`
	Url         string `xml:"Url"`
}
type WxServerEventScanCodeInfo struct {
	ScanType   string `xml:"ScanType"`
	ScanResult string `xml:"ScanResult"`
}
type WxServerEventSendPicsInfo struct { //Event == pic_photo_or_album
	Count   int `xml:"Count"`
	PicList []struct {
		Item struct {
			PicMd5Sum string `xml:"PicMd5Sum"`
		} `xml:"item"`
	} `xml:"PicList"`
}
type WxServerEventSendLocationInfo struct {
	Location_X float64 `xml:"Location_X"`
	Location_Y float64 `xml:"Location_Y"`
	Scale      int     `xml:"Scale"`
	Label      string  `xml:"Label"`
	Poiname    string  `xml:"Poiname"`
}
