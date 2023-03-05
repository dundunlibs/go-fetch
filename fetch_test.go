package gofetch

import (
	"regexp"
	"testing"
)

func TestGetText(t *testing.T) {
	res, err := Fetch("https://example.com")
	if err != nil {
		t.Fatal(err)
	}

	text, err := res.Text()
	if err != nil {
		t.Fatal(err)
	}

	want := regexp.MustCompile("Example Domain")
	if !want.MatchString(text) {
		t.Errorf("Result: %v does not match the pattern of %#q", text, want)
	}
}

func TestGetJSON(t *testing.T) {
	res, err := Fetch("https://api.dictionaryapi.dev/api/v2/entries/en/hello")
	if err != nil {
		t.Fatal(err)
	}

	type Word struct {
		Word string `json:"word"`
	}

	type Body []Word

	var body Body
	err = res.BindJSON(&body)
	if err != nil {
		t.Fatal(err)
	}

	want := "hello"
	if body[0].Word != want {
		t.Errorf("Result: %v does not include %#q", body, want)
	}
}
