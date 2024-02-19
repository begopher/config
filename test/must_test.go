package test

import (
	"github.com/begopher/vars"
	"testing"
)

func Test_Must(t *testing.T) {
	data1 := map[string]string{
		"username": "gopher",
	}
	data2 := map[string]string{
		"port":     "2222",
		"database": "localhost",
	}
	source := vars.Sources(
		vars.Map(data1),
		vars.Map(data2))

	exist, missing := vars.Check(source,
		"username",
		"port",
		"database",
		"password",
		"email")

	if n := len(exist); n != 3 {
		t.Fatalf("3 key must exist got (%d) keys", n)
	}
	if n := len(missing); n != 2 {
		t.Fatalf("2 key must exist got (%d) keys", n)
	}
	for _, key := range exist {
		switch key {
		case "username", "port", "database":
		default:
			t.Errorf("exist has invalid key (%s)", key)
		}
	}
	for _, key := range missing {
		switch key {
		case "email", "password":
		default:
			t.Errorf("missing has invalid key (%s)", key)
		}
	}
}
