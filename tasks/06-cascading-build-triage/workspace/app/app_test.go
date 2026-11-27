package app

import "testing"

func TestSummary(t *testing.T) {
	pro, fee := Summary("u1", "pro")
	if !pro || fee != 20 {
		t.Fatalf("got (%v, %d), want (true, 20)", pro, fee)
	}
}
