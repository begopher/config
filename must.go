package config

func Check(setting Setting, keys ...string) (exist, missing []string) {
	if setting == nil {
		panic("config.Check: setting cannot be nil")
	}
	for _, key := range keys {
		if value := setting.Get(key); value == "" {
			missing = append(missing, key)
			continue
		}
		exist = append(exist, key)
	}
	return
}
