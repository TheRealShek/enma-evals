package parser

import "strings"

func ParsePair(s string) (string, string, bool) {
	k, v, ok := strings.Cut(s, "=")
	if !ok {
		return "", "", false
	}

	return k, v, true
}
