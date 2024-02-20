package test

import (
	"github.com/begopher/config"
	"testing"
)

func Test_Map(t *testing.T) {
	data := map[string]string{
		"username": "gopher",
		"database": "localhost",
		"port":     "1234",
	}
	setting := config.Map(data)
	for key, expected := range data {
		if got := setting.Get(key); got != expected {
			t.Errorf("%v=%s expected value is (%s)", key, got, expected)
		}
	}
}
