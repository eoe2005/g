package gsms

type GSms interface {
	send(msg string) bool
}

var (
	_localsmsmap = map[string]GSms{}
)

func Send(confKey, msg string) bool {
	gs, ok := _localsmsmap[confKey]
	if ok {
		return gs.send(msg)
	}
	return false
}
