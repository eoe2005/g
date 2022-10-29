package gerror

type JsonError struct {
	Code int
	Msg  string
	Data interface{}
}
