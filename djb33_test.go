package cache

import "testing"

// Keys that differ only in their last byte must hash differently; the
// tail handling used to drop the final byte for every remainder length.
func TestDjb33FinalByte(t *testing.T) {
	for _, pair := range [][2]string{{"a", "b"}, {"aaaaa", "aaaab"}, {"abcdefghi", "abcdefghj"},
		{"ab", "ac"}, {"abc", "abd"}, {"abcd", "abce"}, {"abcdefgh", "abcdefgi"}, {"abcdefg", "abcdefh"}, {"abcdef", "abcdeg"}} {
		if djb33(7, pair[0]) == djb33(7, pair[1]) {
			t.Errorf("djb33(%q) == djb33(%q): final byte ignored", pair[0], pair[1])
		}
	}
}
