package middle

var (
	middles = map[string]MiddleHandle{}
)

const (
	WX_SERVER_TOKEN = "wx_server_token"
)

type MiddleHandle func(a ...any) (ret any)

func RegisterMiddle(key string, mid MiddleHandle, a ...any) {
	middles[key] = mid
	lenMid := len(a)
	for i := 0; i < lenMid; i++ {
		k := a[i]
		if i+1 >= lenMid {
			return
		}
		v := a[i+1]
		i += 1
		kk, ok := k.(string)
		val, ok2 := v.(MiddleHandle)
		if !ok || !ok2 {
			return
		}
		middles[kk] = val
	}
}
func Handle(key string, a ...any) (ret any) {
	h, ok := middles[key]
	if !ok {
		return nil
	}
	return h(a...)
}
