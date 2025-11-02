// anagram/anagram_test.go
package anagram_test

import (
	"testing"

	"github.com/abdihaikal/go-anagram/anagram"
)

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		val1     string
		val2     string
		expected bool
	}{
		{"listen", "silent", true},
		{"Debit Card", "Bad Credit", true},
		{"hello", "world", false},
		{"a gentleman", "elegant man", true},
		{"rat", "tar", true},
		{"abc", "abcd", false},
	}

	for _, test := range tests {
		result := anagram.IsAnagram(test.val1, test.val2)
		if result != test.expected {
			t.Errorf("IsAnagram(%q, %q) gagal. Hasil: %t, Diharapkan: %t",
				test.val1, test.val2, result, test.expected)
		}
	}
}
