package gweb

import "github.com/gin-gonic/gin"

type GSession map[string]any

func GetSession(c *gin.Context) GSession {
	v, o := c.Get("session")
	if o {
		return v.(GSession)
	}
	return GSession{}
}
func (s GSession) Set(key string, val any) {
	s[key] = val
}
func (s GSession) Get(key string) any {
	if r, ok := s[key]; ok {
		return r
	}
	return nil
}
func (s GSession) GetInt(key string, defval int64) int64 {
	ret := s.Get(key)
	if ret == nil {
		return defval
	}
	if r, ok := ret.(int64); ok {
		return r
	}
	return defval
}
func (s GSession) GetString(key string, defval string) string {
	ret := s.Get(key)
	if ret == nil {
		return defval
	}
	if r, ok := ret.(string); ok {
		return r
	}
	return defval
}
func (s GSession) GetBool(key string, defval bool) bool {
	ret := s.Get(key)
	if ret == nil {
		return defval
	}
	if r, ok := ret.(bool); ok {
		return r
	}
	return defval
}
func (s GSession) GetFloat(key string, defval float64) float64 {
	ret := s.Get(key)
	if ret == nil {
		return defval
	}
	if r, ok := ret.(float64); ok {
		return r
	}
	return defval
}
func (s GSession) To(obj any) error {
	return nil
}
