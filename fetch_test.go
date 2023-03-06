package gofetch

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetText(t *testing.T) {
	res, err := Fetch("https://nextjs-api-routes-rest.vercel.app/api/users")
	if err != nil {
		t.Fatal(err)
	}

	text, err := res.Text()
	if err != nil {
		t.Fatal(err)
	}

	want := "[{\"id\":1},{\"id\":2},{\"id\":3}]"
	if want != text {
		t.Errorf("Result: %#v does not match %#q", text, want)
	}
}

func TestGetJSON(t *testing.T) {
	res, err := Fetch("https://nextjs-api-routes-rest.vercel.app/api/user/1")
	if err != nil {
		t.Fatal(err)
	}

	type Body struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}

	var body Body
	err = res.BindJSON(&body)
	if err != nil {
		t.Fatal(err)
	}

	want := Body{
		ID:   1,
		Name: "User 1",
	}
	if body.ID != want.ID || body.Name != want.Name {
		t.Errorf("Result: %v does not match %v", body, want)
	}
}

func TestPutJSON(t *testing.T) {
	body := H{
		"id":   1,
		"name": "Foo Bar",
	}

	res, err := Fetch("https://nextjs-api-routes-rest.vercel.app/api/user/1", Options{
		Method: http.MethodPut,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		Body: bytes.NewBuffer([]byte("{\"id\":1,\"name\":\"Foo Bar\"}")),
	})
	if err != nil {
		t.Fatal(err)
	}

	json, err := res.JSON()
	if err != nil {
		t.Fatal(err)
	}

	if json["id"].(float64) != 1 || json["name"] != "Foo Bar" {
		t.Errorf("Result: %v does not match %v", json, body)
	}
}
