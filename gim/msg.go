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
