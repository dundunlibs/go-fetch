package gofetch

import (
	"net/http"
)

type Options struct {
	Method string
}

var DefaultOptions = Options{
	Method: http.MethodGet,
}

func Fetch(url string, opts ...Options) (*Response, error) {
	opt := DefaultOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	req, err := http.NewRequest(opt.Method, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)

	return &Response{
		Response: res,
	}, nil
}
