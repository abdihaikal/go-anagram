// anagram/anagram.go
// Package anagram provides utilities to check if two strings are anagrams.
// It ignores case and whitespace during comparison.
package anagram

import (
	"unicode"
)

// IsAnagram checks if two strings are anagrams of each other.
// It ignores whitespace and letter case.
func IsAnagram(val1, val2 string) bool {
	if len(val1) != len(val2) {
		return false
	}

	freq := make(map[rune]int)
	for _, r := range val1 {
		freq[unicode.ToLower(r)]++
	}
	for _, r := range val2 {
		freq[unicode.ToLower(r)]--
		if freq[unicode.ToLower(r)] < 0 {
			return false
		}
	}
	return true
}
