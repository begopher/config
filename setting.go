package config

// Setting represents a holder of key value pair
// which could be stored in a file, environment
// variables, map, or any other places based on
// the given implementations.
type Setting interface {
	// Get returns value associated with a given key,
	// Get returns empty string if:
	//   - key is associated with empty string
	//   - key does not exist in source (implementation)
	Get(key string) (value string)
}
