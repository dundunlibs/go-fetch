package gofetch

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Options struct {
	Method string
	Header http.Header
	Body   io.Reader
}

var DefaultOptions = Options{
	Method: http.MethodGet,
}

func Fetch(url string, opts ...Options) (*Response, error) {
	opt := DefaultOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	req, err := http.NewRequest(opt.Method, url, opt.Body)
	if err != nil {
		return nil, err
	}

	for k, v := range opt.Header {
		req.Header[k] = v
	}

	res, err := http.DefaultClient.Do(req)

	return &Response{
		Response: res,
	}, nil
}

type H map[string]any

func BodyText(v string) *bytes.Buffer {
	return bytes.NewBufferString(v)
}

func BodyJSON(v any) *bytes.Buffer {
	buf, err := json.Marshal(v)
	if err != nil {
		return nil
	}
	return bytes.NewBuffer(buf)
}
