package gim

import "net"

var (
	_callmap    = map[string]ImCmd{}
	_userComMap = map[uint64][]*ClientCon{}
)

type ImCmd func(msg *GImMessage) error

func AddCmd(cmd string, call ImCmd) {
	_callmap[cmd] = call
}

//登录成功
func LoginSuccess(uid uint64, msg *GImMessage) {
	msg.Client.userId = uid
	msg.UserID = uid
	l, ok := _userComMap[uid]
	if ok {
		l = append(l, msg.Client)
	} else {
		l = []*ClientCon{msg.Client}
	}
	_userComMap[uid] = l
}

//登录失败
func LogoutSuccess(msg *GImMessage) {
	msg.Client.close()
}

//发送消息
func Send(msg *GImMessage, uids ...uint64) {
	for _, uid := range uids {
		l, ok := _userComMap[uid]
		if ok {
			for _, i := range l {
				i.Write(msg)
			}
		}
	}

}
func runCmd(msg *GImMessage) {
	fc, ok := _callmap[msg.Cmd]
	if ok {
		fc(msg)
	}
}

func Run(addr string) {
	l, e := net.Listen("tcp", addr)
	if e != nil {
		panic(e)
	}
	for {
		client, e := l.Accept()
		if e != nil {
			continue
		}
		newClient(client)
	}
}
