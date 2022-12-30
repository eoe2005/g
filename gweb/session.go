package gweb

import (
	"github.com/gin-gonic/gin"
)

type gSession map[string]any

func GetSession(c *gin.Context) *gSession {
	v, o := c.Get("session")
	if o {
		return v.(*gSession)
	}
	ret := &gSession{}
	c.Set("session", ret)
	return ret
}
func (s gSession) Set(key string, val any) {
	s[key] = val
}
func (s gSession) Get(key string) any {
	if r, ok := s[key]; ok {
		return r
	}
	return nil
}
<<<<<<< HEAD
func (s gSession) GetInt(key string, defval int64) int64 {
=======
func (s gSession) GetUInt(key string, defval uint64) uint64 {
>>>>>>> dev
	ret := s.Get(key)
	if ret == nil {
		return defval
	}

	if r, ok := ret.(int); ok {
<<<<<<< HEAD
		return int64(r)
	}
	if r, ok := ret.(int8); ok {
		return int64(r)
	}
	if r, ok := ret.(int16); ok {
		return int64(r)
	}
	if r, ok := ret.(int32); ok {
		return int64(r)
	}
	if r, ok := ret.(int64); ok {
		return int64(r)
	}
	if r, ok := ret.(uint); ok {
		return int64(r)
	}
	if r, ok := ret.(uint8); ok {
		return int64(r)
	}
	if r, ok := ret.(uint16); ok {
		return int64(r)
	}
	if r, ok := ret.(uint32); ok {
		return int64(r)
	}
	if r, ok := ret.(uint64); ok {
		return int64(r)
	}
	if r, ok := ret.(float32); ok {
		return int64(r)
	}
	if r, ok := ret.(float64); ok {
		return int64(r)
=======
		return uint64(r)
	}
	if r, ok := ret.(int8); ok {
		return uint64(r)
	}
	if r, ok := ret.(int16); ok {
		return uint64(r)
	}
	if r, ok := ret.(int32); ok {
		return uint64(r)
	}
	if r, ok := ret.(int64); ok {
		return uint64(r)
	}
	if r, ok := ret.(uint); ok {
		return uint64(r)
	}
	if r, ok := ret.(uint8); ok {
		return uint64(r)
	}
	if r, ok := ret.(uint16); ok {
		return uint64(r)
	}
	if r, ok := ret.(uint32); ok {
		return uint64(r)
	}
	if r, ok := ret.(uint64); ok {
		return r
>>>>>>> dev
	}
	if r, ok := ret.(float32); ok {
		return uint64(r)
	}
	if r, ok := ret.(float64); ok {
		return uint64(r)
	}
	return defval
}
<<<<<<< HEAD
=======
func (s gSession) GetInt(key string, defval int64) int64 {
	ret := s.Get(key)
	if ret == nil {
		return defval
	}

	if r, ok := ret.(int); ok {
		return int64(r)
	}
	if r, ok := ret.(int8); ok {
		return int64(r)
	}
	if r, ok := ret.(int16); ok {
		return int64(r)
	}
	if r, ok := ret.(int32); ok {
		return int64(r)
	}
	if r, ok := ret.(int64); ok {
		return int64(r)
	}
	if r, ok := ret.(uint); ok {
		return int64(r)
	}
	if r, ok := ret.(uint8); ok {
		return int64(r)
	}
	if r, ok := ret.(uint16); ok {
		return int64(r)
	}
	if r, ok := ret.(uint32); ok {
		return int64(r)
	}
	if r, ok := ret.(uint64); ok {
		return int64(r)
	}
	if r, ok := ret.(float32); ok {
		return int64(r)
	}
	if r, ok := ret.(float64); ok {
		return int64(r)
	}
	return defval
}
>>>>>>> dev
func (s gSession) GetString(key string, defval string) string {
	ret := s.Get(key)
	if ret == nil {
		return defval
	}
	if r, ok := ret.(string); ok {
		return r
	}
	return defval
}
func (s gSession) GetBool(key string, defval bool) bool {
	ret := s.Get(key)
	if ret == nil {
		return defval
	}
	if r, ok := ret.(bool); ok {
		return r
	}
	return defval
}
func (s gSession) GetFloat(key string, defval float64) float64 {
	ret := s.Get(key)
	if ret == nil {
		return defval
	}
	if r, ok := ret.(float64); ok {
		return r
	}
	if r, ok := ret.(float32); ok {
		return float64(r)
	}
	return defval
}
func (s gSession) To(obj any) error {
	return nil
}
