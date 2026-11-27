package codec

import "strings"

func EncodeKey(raw string) string {
	return strings.TrimSpace(strings.ToUpper(raw))
}
