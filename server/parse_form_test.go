package server

import (
	"reflect"
	"testing"
)

func TestParseForm(t *testing.T) {
	body := "firstname=John&username=johndoe&password=123"

	got, err := ParseForm([]byte(body))

	if err != nil {
		t.Fatalf("Form(%q) error = %v", body, err)
	}

	wanted := make(map[string]string)

	wanted["firstname"] = "John"
	wanted["username"] = "johndoe"
	wanted["password"] = "123"

	if !reflect.DeepEqual(wanted, got) {
		t.Errorf("wanted = %+v, got = %+v", wanted, got)
	}
}
