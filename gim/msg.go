package gim

import (
	"encoding/json"
)

type GImMessage struct {
	Cmd    string
	Data   []byte
	UserID uint64
	Client *ClientCon
	Obj    any
}

func NewMsg(cmd string, data any) *GImMessage {
	sd, _ := json.Marshal(data)
	return &GImMessage{
		Cmd:  cmd,
		Obj:  data,
		Data: sd,
	}
}

//绑定数据
func (m *GImMessage) BindReq(obj any) error {
	return json.Unmarshal(m.Data, obj)
}
