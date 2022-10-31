package gmail

import (
	"crypto/tls"

	"github.com/eoe2005/g/gconf"
	"gopkg.in/gomail.v2"
)

var (
	_localMailMap = map[string]*gconf.GMailYaml{}
)

func Register(mails []*gconf.GMailYaml) {
	for _, item := range mails {
		_localMailMap[item.Name] = item
	}
}
func New(name string) *gmail {
	msg := gomail.NewMessage()
	if c, ok := _localMailMap[name]; ok {
		msg.SetHeader("From", c.FromName)
		return &gmail{
			conf: c,
			msg:  msg,
		}
	}
	return &gmail{
		msg: msg,
	}
}

type gmail struct {
	conf *gconf.GMailYaml
	msg  *gomail.Message
}

func (m *gmail) From(name string) *gmail {
	m.msg.SetHeader("From", name)
	return m
}
func (m *gmail) To(mails ...string) *gmail {
	m.msg.SetHeader("To", mails...)
	return m
}
func (m *gmail) Cc(mails ...string) *gmail {
	m.msg.SetHeader("Cc", mails...)
	return m
}
func (m *gmail) Bcc(mails ...string) *gmail {
	m.msg.SetHeader("Cc", mails...)
	return m
}
func (m *gmail) Subject(title string) *gmail {
	m.msg.SetHeader("Subject", title)
	return m
}
func (m *gmail) Body(body string) *gmail {
	m.msg.SetBody("text/plain", body)
	return m
}
func (m *gmail) Html(body string) *gmail {
	m.msg.SetBody("text/html", body)
	return m
}
func (m *gmail) AddFile(path string) *gmail {
	m.msg.Attach(path)
	return m
}

// 发送邮件
func (m *gmail) Send() error {
	d := gomail.NewDialer(m.conf.Smtp, m.conf.Prot, m.conf.UserName, m.conf.PassWord)
	if !m.conf.IsSsl {
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}
	return d.DialAndSend(m.msg)
}
