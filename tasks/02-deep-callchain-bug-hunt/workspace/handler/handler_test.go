package handler

import "testing"

func TestHandleRequest(t *testing.T) {
	got, ok := HandleRequest("user_1", "Alice")
	if !ok || got != "Alice" {
		t.Fatalf("HandleRequest = (%q, %v), want (Alice, true)", got, ok)
	}
}
