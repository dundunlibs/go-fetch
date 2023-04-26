package multipart

import (
	"bytes"
	"fmt"
	"io"
	mp "mime/multipart"
)

func NewFormData() *FormData {
	body := &bytes.Buffer{}
	w := mp.NewWriter(body)

	return &FormData{w, body}
}

type FormData struct {
	*mp.Writer
	body *bytes.Buffer
}

func (fd *FormData) Body() *bytes.Buffer {
	fd.Close()
	return fd.body
}

func (fd *FormData) Append(name string, value any, filename ...string) error {
	switch v := value.(type) {
	case *mp.FileHeader:
		return fd.appendFileHeader(name, v)
	case string:
		return fd.WriteField(name, v)
	case []byte:
		w, err := fd.CreateFormField(name)
		if err != nil {
			return err
		}
		_, err = w.Write(v)
		return err
	default:
		return fmt.Errorf("cannot append value type %t", v)
	}
}

func (fd *FormData) appendFileHeader(name string, value *mp.FileHeader) error {
	file, err := value.Open()
	if err != nil {
		return err
	}
	defer file.Close()
	fw, err := fd.CreateFormFile(name, value.Filename)
	if err != nil {
		return err
	}
	_, err = io.Copy(fw, file)
	return err
}
