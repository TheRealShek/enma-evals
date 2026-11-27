package routes

import "testing"

func TestRoutes(t *testing.T) {
	for _, fn := range []func(string) int{StatusA, StatusA2, StatusB, StatusB2, StatusC, StatusC2} {
		if fn("v1") != 200 || fn("v2") != 200 || fn("v3") != 404 {
			t.Fatalf("unexpected status codes")
		}
	}
}
