package test

import (
	"github.com/begopher/vars"
	"testing"
)

func Test_File(t *testing.T) {
	file := "vars.txt"
	source, err := vars.File(file)
	if err != nil {
		t.Fatal(err)
	}
	data := map[string]string{
		"username": "gopher",
		"database": "localhost",
		"port":     "1234",
	}
	for key, expected := range data {
		if got := source.Get(key); got != expected {
			t.Errorf("%v=%s expected value is (%s)", key, got, expected)
		}
	}
}
