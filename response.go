package gofetch

import (
	"encoding/json"
	"io"
	"net/http"
)

type Response struct {
	*http.Response
	body []byte
}

func (res *Response) BodyAsBytes() ([]byte, error) {
	if res.body == nil {
		defer res.Body.Close()
		bytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		res.body = bytes
	}

	return res.body, nil
}

func (res *Response) BindJSON(v any) error {
	body, err := res.BodyAsBytes()
	if err != nil {
		return err
	}

	return json.Unmarshal(body, v)
}

func (res *Response) JSON() (map[string]any, error) {
	var result map[string]any
	err := res.BindJSON(&result)
	return result, err
}

func (res *Response) Text() (string, error) {
	body, err := res.BodyAsBytes()
	if err != nil {
		return "", err
	}
	return string(body), nil
}
