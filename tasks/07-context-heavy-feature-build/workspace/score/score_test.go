package score

import "testing"

func TestNormalizeScore(t *testing.T) {
	cases := [][3]int{
		{5, 0, 0},
		{-3, 10, 0},
		{15, 10, 100},
		{3, 4, 75},
	}

	for _, c := range cases {
		if got := NormalizeScore(c[0], c[1]); got != c[2] {
			t.Fatalf("NormalizeScore(%d, %d) = %d, want %d", c[0], c[1], got, c[2])
		}
	}
}
