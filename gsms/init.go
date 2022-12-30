package gsms

type GSms interface {
	send(mobile, msg string) bool
}

var (
	_localsmsmap = map[string]GSms{}
)

func Send(confKey, mobile, msg string) bool {
	gs, ok := _localsmsmap[confKey]
	if ok {
		return gs.send(mobile, msg)
	}
	return false
}
