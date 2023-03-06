package gutils

import "regexp"

var (
	_regMobile     = regexp.MustCompile("^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$")
	_regEmail      = regexp.MustCompile("^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\\.){1,4}[a-z]{2,4}$")
	_regUserIdCard = regexp.MustCompile("(^\\d{15}$)|(^\\d{18}$)|(^\\d{17}(\\d|x|X)$)")
	_regNickName   = regexp.MustCompile("^[a-z0-9_-]{3,16}$")
	_regPassword   = regexp.MustCompile("^[a-z0-9_-]{6,18}$")
	_regHex        = regexp.MustCompile("^#?([a-f0-9]{6}|[a-f0-9]{3})$")
	_regUrl        = regexp.MustCompile("^(https?://)?([\\da-z\\.-]+)\\.([a-z\\.]{2,6})([\\w \\.-]*)*/?$")
	_regHtml       = regexp.MustCompile("^<([a-z]+)([^<]+)*(?:>(.*)</\\1>|\\s+/>)$")
	_regHanZi      = regexp.MustCompile("^[\\u2E80-\u9FFF]+$")
	_regIP         = regexp.MustCompile("^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$")
	// _regNickName = regexp.MustCompile("")
	// _regNickName = regexp.MustCompile("")
	// _regNickName = regexp.MustCompile("")
)

// 是否是IP
func IsIP(str string) bool {
	return _regIP.MatchString(str)
}

// 是否是汉字
func IsHanZi(str string) bool {
	return _regHanZi.MatchString(str)
}

// 是否是Html
func IsHtml(str string) bool {
	return _regHtml.MatchString(str)
}

// 是否是url
func IsUrl(str string) bool {
	return _regUrl.MatchString(str)
}

// 是否是16进制
func IsHex(str string) bool {
	return _regHex.MatchString(str)
}

// 是否是秘密
func IsPassword(pass string) bool {
	return _regPassword.MatchString(pass)
}

// 是否是昵称
func IsNickName(nickName string) bool {
	return _regNickName.MatchString(nickName)
}

// 是否是手机号
func IsMobile(mobile string) bool {
	return _regMobile.MatchString(mobile)
}

// 是否是油箱
func IsEmail(email string) bool {
	return _regEmail.MatchString(email)
}

// 是否是身份证号
func IsUserIdCard(idCard string) bool {
	return _regUserIdCard.MatchString(idCard)
}
