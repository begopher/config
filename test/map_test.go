package test

import (
	"github.com/begopher/vars"
	"testing"
)

func Test_Map(t *testing.T) {
	data := map[string]string{
		"username": "gopher",
		"database": "localhost",
		"port":     "1234",
	}
	source := vars.Map(data)
	for key, expected := range data {
		if got := source.Get(key); got != expected {
			t.Errorf("%v=%s expected value is (%s)", key, got, expected)
		}
	}
}
