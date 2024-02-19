package vars

import (
	"os"
)
// Environment reads value from environment variables 
func Environment() Source {
	return environment{}
}

type environment struct{}

func (environment) Get(key string) string {
	return os.Getenv(key)
}
