package shard

import (
	"evaltask/keys"
	"testing"
)

func TestActiveHeader(t *testing.T) {
	if got := ActiveHeader(); got != keys.ExpectedHeader() {
		t.Fatalf("ActiveHeader() = %q, want %q", got, keys.ExpectedHeader())
	}
}
