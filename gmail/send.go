package gmail

import "github.com/eoe2005/g/gtemplate"

//发送模板邮件
func SendTemplate(confKey, email, subject, tpl string, data any) error {
	return New(confKey).Cc(email).Subject(subject).Html(gtemplate.Fetch(tpl, data)).Send()
}

//发送html
func SendHtml(confKey, email, subject, html string) error {
	return New(confKey).Cc(email).Subject(subject).Html(html).Send()
}
