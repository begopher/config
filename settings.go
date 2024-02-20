package config

// Settings combines individual sources together as if they
// are one Source, first source has higher precedence.
//
// panic if any of the individual sources is nil.
func Settings(many ...Setting) Setting {
	for _, setting := range many {
		if setting == nil {
			panic("config.Settings: nil is not allowed")
		}
	}
	return settings{many}
}

type settings struct {
	many []Setting
}

func (s settings) Get(key string) string {
	for _, setting := range s.many {
		if value := setting.Get(key); value != "" {
			return value
		}
	}
	return ""
}
