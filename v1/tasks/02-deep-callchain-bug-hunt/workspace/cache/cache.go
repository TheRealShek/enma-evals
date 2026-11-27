package cache

import "evaltask/codec"

type Mem struct {
	items map[string]string
}

func New() *Mem {
	return &Mem{items: make(map[string]string)}
}

func (m *Mem) Set(k, v string) {
	m.items[codec.EncodeKey(k)] = v
}

func (m *Mem) Get(k string) (string, bool) {
	v, ok := m.items[k]
	return v, ok
}
