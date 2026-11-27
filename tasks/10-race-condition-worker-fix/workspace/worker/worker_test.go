package worker

import (
	"sync"
	"testing"
)

func TestRegistryConcurrent(t *testing.T) {
	reg := NewRegistry()
	var wg sync.WaitGroup

	for range 20 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 50 {
				reg.Inc("hits")
			}
		}()
	}

	wg.Wait()

	if got := reg.Get("hits"); got != 1000 {
		t.Fatalf("Get(hits) = %d, want 1000", got)
	}
}
