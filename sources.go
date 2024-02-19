package vars

// Sources combines individual sources together and acts
// as if it is one Source, first source has higher precedence.
//
// panic if any of the individual sources is nil.
func Sources(many ...Source) Source {
	for _, source := range many {
		if source == nil {
			panic("vars.Sources: nil argument is not allowed")
		}
	}
	return sources{many}
}

type sources struct {
	many []Source
}

func (s sources) Get(key string) string {
	for _, source := range s.many {
		if value := source.Get(key); value != "" {
			return value
		}
	}
	return ""
}
