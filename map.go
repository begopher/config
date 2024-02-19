package vars

// Map reads value from given map.
func Map(source map[string]string) Source {
	if source == nil {
		source = make(map[string]string)
	}
	return _map{source}
}

type _map struct {
	vars map[string]string
}

func (m _map) Get(key string) string {
	return m.vars[key]
}
