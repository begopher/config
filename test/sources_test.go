package test

import (
	"github.com/begopher/vars"
	"testing"
)

func Test_Sources(t *testing.T) {
	top := map[string]string{
		"port": "1111",
	}
	mid := map[string]string{
		"port":     "2222",
		"database": "localhost2",
	}
	low := map[string]string{
		"port":     "3333",
		"database": "localhost3",
		"username": "gopher3",
	}
	source := vars.Sources(
		vars.Map(top),
		vars.Map(mid),
		vars.Map(low))

	expected := "1111"
	if got := source.Get("port"); got != expected {
		t.Errorf("expected port is (%s) got (%s)", expected, got)
	}
	expected = "localhost2"
	if got := source.Get("database"); got != expected {
		t.Errorf("expected database is (%s) got (%s)", expected, got)
	}
	expected = "gopher3"
	if got := source.Get("username"); got != expected {
		t.Errorf("expected username is (%s) got (%s)", expected, got)
	}
	expected = ""
	if got := source.Get("password"); got != expected {
		t.Errorf(`expected password is empty string () got (%s)`, got)
	}
}
