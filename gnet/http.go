package gnet

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type HttResult struct {
	Code int
	Body []byte
	Err  error
}

func Get(ctx context.Context, url string, timeout int, headers map[string]string) *HttResult {
	return doSend(ctx, http.MethodGet, url, timeout, headers, nil)
}
func Post(ctx context.Context, url, data string, timeout int, headers map[string]string) *HttResult {
	return doSend(ctx, http.MethodPost, url, timeout, headers, strings.NewReader(data))
}
func PostForm(ctx context.Context, url string, data url.Values, timeout int, headers map[string]string) *HttResult {
	if headers == nil {
		headers = map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	}
	return doSend(ctx, http.MethodPost, url, timeout, headers, strings.NewReader(data.Encode()))
}

func doSend(ctx context.Context, method, url string, timeout int, headers map[string]string, body io.Reader) *HttResult {
	ret := &HttResult{
		Code: 9999,
		Body: nil,
	}
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		ret.Err = err
		return ret
	}
	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	client := &http.Client{}
	if timeout > 0 {
		client.Timeout = time.Duration(timeout) * time.Millisecond
	}
	rep, err := client.Do(req)
	if err != nil {
		ret.Err = err
		return ret
	}
	defer rep.Body.Close()
	r, e := ioutil.ReadAll(rep.Body)
	ret.Body = r
	ret.Err = e
	return ret

}
