package test

import (
	"github.com/begopher/config"
	"testing"
)

func Test_File(t *testing.T) {
	file := "vars.txt"
	setting, err := config.File(file)
	if err != nil {
		t.Fatal(err)
	}
	data := map[string]string{
		"username": "gopher",
		"database": "localhost",
		"port":     "1234",
	}
	for key, expected := range data {
		if got := setting.Get(key); got != expected {
			t.Errorf("%v=%s expected value is (%s)", key, got, expected)
		}
	}
}
