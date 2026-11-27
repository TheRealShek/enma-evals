package tag

import "testing"

func TestFormatTag(t *testing.T) {
	if got := FormatTag("  GoLang "); got != "#golang" {
		t.Fatalf("FormatTag = %q, want #golang", got)
	}
}
