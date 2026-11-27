package parser

import "testing"

func TestParsePair(t *testing.T) {
	k, v, ok := ParsePair(" host = localhost ")
	if !ok || k != "host" || v != "localhost" {
		t.Fatalf("got (%q, %q, %v), want (host, localhost, true)", k, v, ok)
	}
}
