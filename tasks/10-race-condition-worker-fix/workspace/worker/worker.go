package worker

import "sync"

type Registry struct {
	mu     sync.Mutex
	counts map[string]int
}

func NewRegistry() *Registry {
	return &Registry{counts: make(map[string]int)}
}

func (r *Registry) Inc(key string) {
	r.counts[key]++
}

func (r *Registry) Get(key string) int {
	return r.counts[key]
}
