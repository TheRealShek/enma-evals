package app

import "testing"

func TestReadSetting(t *testing.T) {
	if got := ReadSetting(map[string]string{"mode": "fast"}, "mode"); got != "fast" {
		t.Fatalf("got %q, want fast", got)
	}
}
