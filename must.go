package vars

func Check(source Source, keys ...string) (exist, missing []string) {
	if source == nil {
		panic("vars.Must: source cannot be nil")
	}
	for _, key := range keys {
		if value := source.Get(key); value == "" {
			missing = append(missing, key)
			continue
		}
		exist = append(exist, key)
	}
	return
}
