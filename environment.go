package config

import (
	"os"
)
// Environment variables as source of setting
func Environment() Setting {
	return environment{}
}

type environment struct{}

func (environment) Get(key string) string {
	return os.Getenv(key)
}
