package server

import "testing"

func TestRouterLimit(t *testing.T) {
	r := NewRouter(2)
	if r.Handle() != 200 || r.Handle() != 200 || r.Handle() != 429 {
		t.Fatalf("expected 200, 200, 429")
	}
}
