package test

import (
	"github.com/begopher/config"
	"os"
	"testing"
)

func Test_Enviorment(t *testing.T) {
	data := map[string]string{
		"vars_test_username": "gopher",
		"vars_test_database": "localhost",
		"vars_test_port":     "1234",
	}
	for key, value := range data {
		if err := os.Setenv(key, value); err != nil {
			t.Fatal(err)
		}
		defer os.Unsetenv(key)
	}
	setting := config.Environment()
	for key, expected := range data {
		if got := setting.Get(key); got != expected {
			t.Errorf("%v=%s expected value is (%s)", key, got, expected)
		}
	}
}
