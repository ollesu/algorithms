package algorithms

import "strings"

func reverse(s string) string {
	var reversed string
	for i := len(s) - 1; i >= 0; i-- {
		reversed = reversed + string(s[i])
	}
	return reversed
}

// more efficient way in Go using strings.Builder (similar to bytes.Buffer):
func reverseWithBuilder(s string) string {
	var sb strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteString(string(s[i]))
		// or sb.WriteByte(s[i])
	}
	return sb.String()
}

// using runes to accommodate
// rather than dealing with string characters as bytes, we deal with them as runes to accommodate non-ASCII chars
func reverseRunes(s string) string {
	var reversed string
	for _, r := range s {
		reversed = string(r) + reversed
	}
	return reversed
}
