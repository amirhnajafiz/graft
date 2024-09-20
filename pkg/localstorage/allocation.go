package localstorage

type memory map[string][]interface{}

func (m memory) Put(key string, value interface{}) {
	if _, ok := m[key]; !ok {
		m[key] = make([]interface{}, 0)
	}

	m[key] = append(m[key], value)
}

func (m memory) Fetch(key string) interface{} {
	return m[key]
}
