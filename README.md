# go-fetch
JS-like Fetch API for Go `net/http`

## Installation
```bash
go get github.com/dundunlabs/go-fetch
```

## Usage
```go
package example

import (
    "fmt"
    "net/http"
    "github.com/dundunlabs/go-fetch"
)

func GetExample() {
    res, err := gofetch.Fetch("https://example.com")
    if err != nil {
        fmt.Println("failed to get data: ", err)
	}

	text, err := res.Text()
	if err != nil {
		fmt.Println("failed to read body as text: ", err)
	}

    fmt.Println(text)
}

func PostExample() {
    res, err := gofetch.Fetch("https://example.com", gofetch.Options{
        Method: http.MethodPost,
        Header: http.Header{
            "Content-Type": []string{"application/json"},
        },
        Body: gofetch.BodyJSON(gofetch.H{
            "foo":   "Bar",
        }),
    })
	if err != nil {
        fmt.Println("failed to post data: ", err)
	}

	json, err := res.JSON()
	if err != nil {
		fmt.Println("failed to read body as json: ", err)
	}

    fmt.Println(json)
}

```