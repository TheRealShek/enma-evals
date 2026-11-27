package memory

type MemStore struct {
	Data map[string]string
}

func (m *MemStore) Get(key string) string {
	return m.Data[key]
}
