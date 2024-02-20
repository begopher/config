package test

import (
	"github.com/begopher/config"
	"testing"
)

func Test_Settings(t *testing.T) {
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
	setting := config.Settings(
		config.Map(top),
		config.Map(mid),
		config.Map(low))

	expected := "1111"
	if got := setting.Get("port"); got != expected {
		t.Errorf("expected port is (%s) got (%s)", expected, got)
	}
	expected = "localhost2"
	if got := setting.Get("database"); got != expected {
		t.Errorf("expected database is (%s) got (%s)", expected, got)
	}
	expected = "gopher3"
	if got := setting.Get("username"); got != expected {
		t.Errorf("expected username is (%s) got (%s)", expected, got)
	}
	expected = ""
	if got := setting.Get("password"); got != expected {
		t.Errorf(`expected password is empty string () got (%s)`, got)
	}
}
