package service

import "testing"

type mockStore map[string]string

func (m mockStore) Get(key string) string {
	return m[key]
}

func TestLookupOrDefault(t *testing.T) {
	s := mockStore{"env": "prod", "empty": ""}
	if got := LookupOrDefault(s, "env", "dev"); got != "prod" {
		t.Fatalf("got %q, want prod", got)
	}
	if got := LookupOrDefault(s, "missing", "dev"); got != "dev" {
		t.Fatalf("got %q, want dev", got)
	}
}
