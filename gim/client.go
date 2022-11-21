package gim

import (
	"net"
	"time"
)

type ClientCon struct {
	con    net.Conn
	msg    chan *GImMessage
	userId uint64
	isRun  bool
}

func newClient(con net.Conn) {
	ret := &ClientCon{
		con:   con,
		isRun: true,
		msg:   make(chan *GImMessage, 10),
	}
	go ret.loopRead()
	go ret.loopWrite()
}

//是否登录
func (c *ClientCon) IsLogin() bool {
	return c.userId > 0
}

//循环读取数据
func (c *ClientCon) loopRead() {
	for c.isRun {
		msg := c.readMsg()
		if msg != nil {
			runCmd(msg)
		}
	}
}
func (c *ClientCon) readMsg() *GImMessage {

	bl := make([]byte, 1)
	n, e := c.con.Read(bl)
	if n == 0 || e != nil {
		return nil
	}
	nl := int(bl[0]) + 4
	cmd := make([]byte, nl)
	n, e = c.con.Read(cmd)
	if n != nl || e != nil {
		return nil
	}
	ml := int(cmd[nl-1]<<24) + (int(cmd[nl-2]) << 16) + (int(cmd[nl-3]) << 8) + (int(cmd[nl-4]))
	data := make([]byte, ml)

	n, e = c.con.Read(data)
	if n != ml || e != nil {
		return nil
	}
	return &GImMessage{
		Cmd:    string(cmd[:bl[0]]),
		Data:   data,
		UserID: c.userId,
		Client: c,
	}
}

//写入数据
func (c *ClientCon) Write(msg *GImMessage) {
	c.msg <- msg
}

//循环推送数据
func (c *ClientCon) loopWrite() {
	for c.isRun {
		select {
		case m := <-c.msg:
			cmd := []byte(m.Cmd)
			data := make([]byte, 1)
			data[0] = byte(len(cmd))
			data = append(data, cmd...)
			ml := len(m.Data)
			data = append(data, byte(ml>>24), byte(ml>>16), byte(ml>>8))
			data = append(data, m.Data...)
			_, e := c.con.Write(data)
			if e != nil {
				c.close()
			}
		default:
			time.Sleep(time.Millisecond * 50)
		}
	}
}
func (c *ClientCon) close() {
	c.isRun = false
}
