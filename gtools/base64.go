package gtools

import "encoding/base64"

func Base64Encode(src string) string {
	return base64.RawStdEncoding.EncodeToString([]byte(src))
}
func Base64UrlEncode(src string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(src))
}

func Base64Decode(src string) string {
	ret, e := base64.RawStdEncoding.DecodeString(src)
	if e != nil {
		return ""
	}
	return string(ret)
}
func Base64UrlDecode(src string) string {
	ret, e := base64.RawURLEncoding.DecodeString(src)
	if e != nil {
		return ""
	}
	return string(ret)
}
