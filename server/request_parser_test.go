package server

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	input := "POST /users HTTP/1.1\r\n" +
		"Content-Type: application/x-www-form-urlencoded\r\n" +
		"Content-Length: 14\r\n" +
		"\r\n" +
		"firstname=John"

	r := bufio.NewReader(strings.NewReader(input))

	got, err := Parse(r)

	if err != nil {
		t.Fatalf("Parse(%q) error = %v", input, err)
	}

	want := &Request{
		method: "POST",
		path:   "/users",
		header: map[string]string{
			"content-type":   "application/x-www-form-urlencoded",
			"content-length": "14",
		},
		body: []byte("firstname=John"),
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
